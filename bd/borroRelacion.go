package bd

import (
	"context"
	"time"

	"github.com/lutrueba/microblog/models"
)

/*BorroRelacion borra la relación en la bd */
func BorroRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microblog")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, rel)
	if err != nil {
		return false, err
	}
	return true, nil
}
