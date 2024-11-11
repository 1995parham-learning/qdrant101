package main

import (
	"context"
	"log"

	"github.com/qdrant/go-client/qdrant"
)

func main() {
	// The Go client uses Qdrant's gRPC interface
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		log.Fatalf("qdrant connection failed %s", err)
	}

	if err := client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: "eli_collection",
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     4,
			Distance: qdrant.Distance_Dot,
		}),
	}); err != nil {
		log.Fatalf("failed to create collection %s", err)
	}

	operationInfo, err := client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: "eli_collection",
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(1),
				Vectors: qdrant.NewVectors(0.05, 0.61, 0.76, 0.74),
				Payload: qdrant.NewValueMap(map[string]any{"city": "Berlin"}),
			},
			{
				Id:      qdrant.NewIDNum(2),
				Vectors: qdrant.NewVectors(0.19, 0.81, 0.75, 0.11),
				Payload: qdrant.NewValueMap(map[string]any{"city": "London"}),
			},
			{
				Id:      qdrant.NewIDNum(3),
				Vectors: qdrant.NewVectors(0.36, 0.55, 0.47, 0.94),
				Payload: qdrant.NewValueMap(map[string]any{"city": "Moscow"}),
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to do upsert %s", err)
	}

	log.Println(operationInfo)
}
