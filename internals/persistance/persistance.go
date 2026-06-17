package persistance

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rishav2006/redis-clone/internals/store"
)

type Snapshot struct {
	Data       map[string]string
	Expiration map[string]time.Time
}

func SaveSnapshot() {
	var snapshot = Snapshot{
		Data:       store.DB.Data,
		Expiration: store.DB.Expiration,
	}
	jsonData, err := json.Marshal(snapshot)
	if err != nil {
		fmt.Println("Some error converting to json", err.Error())
		return
	}
	// filePath := "../../dump.json"
	err = os.WriteFile("dump.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error printing the output onto the file")
	}
}

func LoadSnapshot() {
	fileData, err := os.ReadFile("dump.json")
	if err != nil {
		fmt.Println("No snapshot found...Starting fresh database", err.Error())
		return
	}

	var snapshot Snapshot
	err = json.Unmarshal(fileData, &snapshot)

	if err != nil {
		fmt.Println("Error while unmarshalling json file content", err.Error())
		return
	}

	store.DB.Data = snapshot.Data
	store.DB.Expiration = snapshot.Expiration
	fmt.Println("All snapshots recovered")
}
