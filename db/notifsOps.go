package db

// RegisterForNotifs adds a Channel to the notify table
func (DB *Database) RegisterForNotifs(ChannelID string) error {
	queryString := "INSERT INTO notify (ChannelID) VALUES (:ChannelID)"

	_, err := DB.db.NamedExec(
		queryString,
		map[string]interface{}{
			"ChannelID": ChannelID,
		})
	return err
}

// GetRegisteredChannels returns all ChannelIDs from
// the notify table
func (DB *Database) GetRegisteredChannels() ([]string, error) {
	var channels []string
	err := DB.db.Select(&channels, "SELECT ChannelID from notify")

	return channels, err
}
