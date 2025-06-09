package constants

const (
	Schema string = `
			CREATE TABLE IF NOT EXISTS scheduler (
    			id INTEGER PRIMARY KEY AUTOINCREMENT,
    			date CHAR(8) NOT NULL DEFAULT "",
    			title VARCHAR(255) NOT NULL DEFAULT "",
    			comment TEXT,
    			repeat VARCHAR(128) NOT NULL DEFAULT ""
			);

			CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);
`
)
