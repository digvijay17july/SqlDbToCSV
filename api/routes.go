package api

import (
	"archive/zip"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func addRoutes(r *gin.Engine, store *Store) {
	api := r.Group("/api")
	api.GET("/loadData", func(ctx *gin.Context) {
		log.Info().Msg("Starting the DB Set up")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		ctx.Writer.Header().Set("Content-type", "application/octet-stream")
		ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=tablesData.zip")
		// filePtr, err := os.Open("tablesData.zip")
		// if err != nil {
		// 	log.Error().Err(err)
		// }
		MyZipWriter := zip.NewWriter(ctx.Writer)
		tables := GetTableNames()

		for _, tableName := range tables {

			page := 1
			count := CheckCount(tableName)
			totalPages := count / dbConfig.RowsPerFile
			for i := 0; i < totalPages; i++ {

				data := ListRecords(page, tableName, count)
				// b := &bytes.Buffer{} // creates IO Writer
				writer, err := MyZipWriter.Create(tableName + "_" + strconv.Itoa(page) + ".csv")
				if err != nil {
					log.Error().Err(err)
				}

				wr := csv.NewWriter(writer)
				// creates a csv writer that uses the io buffer.

				// for _, row := range data { // make a loop for 100 rows just for testing purposes

				// }

				wr.WriteAll(data) // converts array of string to comma seperated values for 1 row.
				// _, err = writer.Write([]byte("Sample text"))
				// if err != nil {
				// 	fmt.Println(err)
				// }
				wr.Flush()
				// writes the csv writer data to  the buffered data io writer(b(bytes.buffer))
				page = page + 1
			}
		}
		MyZipWriter.Flush()
		err := MyZipWriter.Close()
		if err != nil {
			log.Error().Err(err)
		}

		// filePtr.Close()

		// ctx.Data(http.StatusOK, "text/zip", MyZipWrite)
		// ctx.FileAttachment(filePtr.)

	})
	api.POST("/loadConfig", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		var dbConfig DBConfig
		//validate the request body
		if err := ctx.ShouldBindJSON(&dbConfig); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		Connect(dbConfig)
		ctx.JSON(http.StatusOK, dbConfig)
	})
	api.DELETE("/todos/:id", func(ctx *gin.Context) {})
	api.PUT("/todos/:id", func(ctx *gin.Context) {})
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
