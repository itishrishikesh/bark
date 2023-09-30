package tests

import (
	"github.com/techrail/bark/client"
	"github.com/techrail/bark/models"
	"testing"
)

func Test_requester(t *testing.T) {
	logClient := client.NewClient("http://localhost:8080/", "INFO", "ServicName", "localRun")

	// Print with formatter
	logClient.Error("L#1L3WBF - Anime: Naruto")
	logClient.Info("L#1L3WBF Anime: One Piece")
	logClient.Debug(" -- Anime: Bleach")
	logClient.Warn("-- Anime: AOT")

	// Print without formatter
	logClient.Errorf("L#1L3WBF - Anime: %s", "Full Metal Alchemist")
	logClient.Infof("L#1L3WBF - Anime: %s", "Tokyo Ghoul")
	logClient.Warnf("L#1L3WBF - Anime: %s", "")
	logClient.Debugf("L#1L3WBF - I want to print something! %s", "weirdString")

	// Multiple Logs
	var logs []models.BarkLog
	logs = make([]models.BarkLog, 3)
	logs[0] = models.BarkLog{Message: "someMessage"}
	logs[1] = models.BarkLog{Message: "someMessage"}
	logs[2] = models.BarkLog{Message: "someMessage"}
	_, _ = client.PostLogArray("http://localhost:8080/insertMultiple", logs)
}
