use chrono::Local;
use std::fs::OpenOptions;
use std::io::Write;

pub fn log_action(action: &str) {
    let now = Local::now();
    let timestamp = now.format("%Y-%m-%d %H:%M:%S").to_string();
    let log_entry = format!("[{}] {}
", timestamp, action);

    if let Ok(mut file) = OpenOptions::new()
        .create(true)
        .append(true)
        .open("fileinbreach.log")
    {
        let _ = file.write_all(log_entry.as_bytes());
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::fs;

    #[test]
    fn test_log_action() {
        let log_file = "fileinbreach.log";
        // Remove existing log if any
        let _ = fs::remove_file(log_file);

        log_action("Test action");
        
        let content = fs::read_to_string(log_file).unwrap();
        assert!(content.contains("Test action"));
        
        // Cleanup
        let _ = fs::remove_file(log_file);
    }
}
