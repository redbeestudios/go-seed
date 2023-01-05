package sql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"

	_ "github.com/denisenkom/go-mssqldb"
)

var _ out.PokemonRepository = (*SqlPokemonRestAdapter)(nil)

const POKEMON_BY_NAME_TABLE = "pokemon_with_name"

type SqlPokemonRestAdapter struct {
	database   *sql.DB
	repository out.PokemonRepository
}

func (a *SqlPokemonRestAdapter) SavePokemon(ctx context.Context, pokemon *pokemon.Pokemon) error {
	query := "INSERT INTO dbo.Pokemon (DexNumber, Name, Type1, Type2) VALUES ('" + strconv.Itoa(pokemon.Id()) + "', '" + pokemon.Name() + "', '" + pokemon.Type().String() + "', '" + pokemon.SecondaryType().String() + "');"
	_, err := a.database.Exec(query)

	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
		return err
	}
	return err
}

func NewSqlPokemonRestAdapter(
	repository out.PokemonRepository,
) *SqlPokemonRestAdapter {

	query := url.Values{}
	query.Add("app name", "MyAppName")
	query.Add("database", "Pokemon-Test")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", "Rodrigo29!"),
		Host:   fmt.Sprintf("%s:%d", "localhost", 1433),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())

	if err != nil {
		log.Fatal(fmt.Errorf("could not initialize mssql server connection"))
		return nil
	}

	return &SqlPokemonRestAdapter{
		database:   db,
		repository: repository,
	}
}

func (a *SqlPokemonRestAdapter) GetByName(
	ctx context.Context,
	name string,
) (*pokemon.Pokemon, error) {

	stmt, err := a.database.Prepare("select * from dbo.Pokemon")

	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var pkmnName string
	var type1 string
	var type2 *string
	var dexNumber int
	var id int
	err = row.Scan(&pkmnName, &type1, &type2, &dexNumber, &id)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("pkmnName:%s\n", pkmnName)
	fmt.Printf("type1:%s\n", type1)
	fmt.Printf("type2:%s\n", &type2)
	fmt.Printf("dexNumber:%d\n", dexNumber)
	fmt.Printf("id:%d\n", id)

	fmt.Printf("bye\n")
	return nil, nil
}
