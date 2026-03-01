use sysinfo::{System, Pid};
use std::process::Command;
use windows::Win32::System::Threading::{OpenProcess, TerminateProcess, PROCESS_TERMINATE};
use windows::Win32::Foundation::{CloseHandle, HANDLE};

pub fn is_critical_process(name: &str) -> bool {
    let critical = [
        "System", "Idle", "smss.exe", "csrss.exe", "wininit.exe", 
        "services.exe", "lsass.exe", "lsm.exe", "svchost.exe"
    ];
    critical.iter().any(|&c| c.eq_ignore_ascii_case(name))
}

pub fn kill_process(pid: u32, _force: bool) -> Result<(), String> {
    unsafe {
        let handle_res = OpenProcess(PROCESS_TERMINATE, false, pid);
        match handle_res {
            Ok(handle) => {
                let res = TerminateProcess(handle, 1);
                let _ = CloseHandle(handle);
                if res.is_ok() {
                    Ok(())
                } else {
                    Err(format!("Failed to terminate process {}: {:?}", pid, res.err()))
                }
            }
            Err(e) => {
                // Try fallback with taskkill if OpenProcess fails (e.g. permission issue)
                let output = Command::new("taskkill")
                    .arg("/F")
                    .arg("/PID")
                    .arg(pid.to_string())
                    .output();
                
                match output {
                    Ok(out) if out.status.success() => Ok(()),
                    Ok(out) => Err(format!("Taskkill failed for PID {}: {}", pid, String::from_utf8_lossy(&out.stderr))),
                    Err(_) => Err(format!("OpenProcess failed for PID {}: {:?}", pid, e)),
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::process::Stdio;

    #[test]
    fn test_kill_process() {
        // Spawn a dummy process (e.g., a long sleep)
        let mut child = Command::new("powershell")
            .arg("-Command")
            .arg("Start-Sleep -Seconds 10")
            .stdout(Stdio::null())
            .stderr(Stdio::null())
            .spawn()
            .expect("Failed to spawn dummy process");

        let pid = child.id();
        
        // Wait a bit for it to start
        std::thread::sleep(std::time::Duration::from_millis(100));

        // Try to kill it
        let result = kill_process(pid, true);
        
        assert!(result.is_ok(), "Failed to kill process {}: {:?}", pid, result.err());
        
        // Wait for it to die
        std::thread::sleep(std::time::Duration::from_millis(500));

        // Verify it's actually gone
        let mut sys = System::new_all();
        sys.refresh_all();
        assert!(sys.process(Pid::from(pid as usize)).is_none(), "Process {} still exists", pid);
    }
    
    #[test]
    fn test_is_critical_process() {
        assert!(is_critical_process("System"));
        assert!(is_critical_process("svchost.exe"));
        assert!(is_critical_process("SVCHOST.EXE"));
        assert!(!is_critical_process("notepad.exe"));
    }
}
