package main

import (
    "bytes"
    "fmt"
    "net/http"
    "os"
    "io/ioutil"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/static"
    "github.com/plotly/graphpiper/plotly"
)

func main() {
    r := gin.Default()

    // Serve static files (CSS, JS, etc.)
    r.Use(static.Serve("/", static.LocalFile("./static", true))

    // Define a route for the root page
    r.GET("/", func(c *gin.Context) {
        // Read the data from the Excel file
        dataURL := "https://github.com/Brutosippon/dados_cv/blob/main/db_PIB_stats_capeverde.xlsx?raw=true"
        dataResponse, err := http.Get(dataURL)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer dataResponse.Body.Close()
        data, err := ioutil.ReadAll(dataResponse.Body)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Create a Plotly figure for the bar chart
        fig1 := plotly.NewFigure()
        fig1.Bar(data, "ano", "Produto_Interno_Bruto", "Inflacao_Media_Anual")
        fig1.SetBarmode("group")

        // Create a Plotly figure for the scatter plot
        fig2 := plotly.NewFigure()
        fig2.Scatter(data, "ano", "Produto_Interno_Bruto", "Life_expectancy_at_birth_total_(years)", "PIB_per_capita(US)", "Stock_da_Divida_Externa_")
        fig2.SetLogX(true)
        fig2.SetSizeMax(60)

        // Render the HTML page with the figures
        c.HTML(http.StatusOK, "index.html", gin.H{
            "fig1": fig1.Render(),
            "fig2": fig2.Render(),
        })
    })

    // Run the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
