package bd

import (
	"context"
	"time"
	"github.com/lutrueba/microblog/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya está en la BD*/
func ChequeoYaExisteUsuario(email string)(models.Usuario,bool,string) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("microblog")
	col := db.Collection("usuarios")
	
	condicion := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx,condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado,false,ID
	}
	return resultado,true,ID

}