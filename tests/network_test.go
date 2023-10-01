package tests

import (
	"github.com/techrail/bark/client"
	"testing"
)

func Test_requester(t *testing.T) {
	logClient := client.NewClient("http://localhost:8080/", "INFO", "ServicName", "localRun")

	// Print with formatter
	logClient.Error("E#1L3WBF - Anime: Naruto")
	logClient.Info("I#1L3WBF - Anime: One Piece")
	logClient.Debug("D#1L3WBF - Anime: Bleach")
	logClient.Warn("W#1L3WBF - Anime: AOT")

	// Print without formatter
	logClient.Errorf("E#1L3WBF - Anime: %s", "Full Metal Alchemist")
	logClient.Infof("I#1L3WBF - Anime: %s", "Tokyo Ghoul")
	logClient.Warnf("W#1L3WBF - Anime: %s", "Haikyuu")
	logClient.Debugf("D#1L3WBF - Anime: %s", "Demon Slayer")

	// Multiple Logs
	//var logs []models.BarkLog
	//logs = make([]models.BarkLog, 3)
	//logs[0] = models.BarkLog{Message: "someMessage"}
	//logs[1] = models.BarkLog{Message: "someMessage"}
	//logs[2] = models.BarkLog{Message: "someMessage"}
	//_, _ = client.PostLogArray("http://localhost:8080/insertMultiple", logs)
}
