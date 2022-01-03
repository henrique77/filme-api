package servicos

import (
	"context"
	"encoding/json"
	"filme-api/mongodb"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Estrutura para persistir no banco de dados
type Movie struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Year  string `json:"year,omitempty"`
}

//Constante do banco de dados e da colecão
const dbName = "filmesdb"
const collectionName = "filmes"

/*	BuscarMovie - Lista todos os filmes salvo no banco de dados
	- ou lista algum filme usando o ID como parâmetro */
func BuscarMovie(c *fiber.Ctx) error {
	//Conecta com o mangoDB e a coleção
	collection, err := mongodb.GetMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500)
		return err
	}

	var filter bson.M = bson.M{}

	//Quando informado o ID, lista o filme especifico
	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500)
		return err
	}

	cur.All(context.Background(), &results)

	//Quando ID informado não exite no banco de dados
	if results == nil {
		c.SendStatus(404)
		return err
	}

	json, _ := json.Marshal(results)
	c.Send(json)
	return err

}

// AddMovie - Insere um filme no banco de dados
func AddMovie(c *fiber.Ctx) error {
	//Conecta com o mangoDB e a coleção
	collection, err := mongodb.GetMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500)
		return err
	}

	var movie Movie
	json.Unmarshal([]byte(c.Body()), &movie)

	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)
	return err

}

//Deleta o filme com o ID informado
func DelMovie(c *fiber.Ctx) error {
	//Conecta com o mangoDB e a coleção
	collection, err := mongodb.GetMongoDbCollection(dbName, collectionName)

	if err != nil {
		c.Status(500)
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500)
		return err
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
	return err

}
