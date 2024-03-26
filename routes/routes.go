package routes

import (
	"database/sql"
	"net/http"

	"github.com/vitorwhois/Carnavou/handlers"
)

func InitRoutes(db *sql.DB) {
	indexHandler := handlers.NewIndexHandler()
	minhalistaHandler := handlers.NewMinhalistaHandler()
	sobreHandler := handlers.NewSobreHandler()
	BlocosPorDataHandler := handlers.NewBlocosPorDataHandler(db)
	BlocosPorNomeHandler := handlers.NewBlocosPorNomeHandler(db)
	BlocosPorIDHandler := handlers.NewBlocosPorIDHandler(db)
	BlocosPorDataSubprefeituraHandler := handlers.NewBlocosPorDataSubprefeituraHandler(db)
	BlocosPorSubprefeituraHandler := handlers.NewBlocosPorSubprefeituraHandler(db)

	http.HandleFunc("/", indexHandler.Handle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/minhalista", minhalistaHandler.Handle)
	http.HandleFunc("/sobre", sobreHandler.Handle)

	http.Handle("/blocos/datas", BlocosPorDataHandler)
	http.Handle("/blocos/nomes", BlocosPorNomeHandler)
	http.Handle("/blocos/ids", BlocosPorIDHandler)
	http.Handle("/blocos/filtros", BlocosPorDataSubprefeituraHandler)
	http.Handle("/blocos/subprefeituras", BlocosPorSubprefeituraHandler)
}
