package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type datos struct {
	op1 string
	op2 string
}

type datosCSV struct {
	op1       string
	filas     int
	columnas  int
	nombre    string
	valMinimo int
	valMaximo int
}

func createHTMLgraphs() { //Funcion para crear los archivos HTML de la graficas
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		Theme: types.ThemeChalk,
	}),
		charts.WithTitleOpts(opts.Title{
			Title:    "GRAFICO DE BARRAS",
			Subtitle: "Con datos generados aleatoriamente",
		}))
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	a, _ := os.Create("bar.html")
	bar.Render(a)

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeChalk,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "GRAFICO DE LINEAS",
			Subtitle: "Con datos generados aleatoriamente",
		}),
	)
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	b, _ := os.Create("line.html")
	line.Render(b)

	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		Theme: types.ThemeChalk,
	}),
		charts.WithTitleOpts(
			opts.Title{
				Title:    "GRAFICA DE PASTEL",
				Subtitle: "Con datos generados aleatoriamente",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		generatePieItems()).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: 200,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				},
			),
		)
	c, _ := os.Create("pie.html")
	pie.Render(c)

}

//-----------------------------------------------------------------------------//

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func archivoCSVBarItems(op2 string, fila int) []opts.BarData {
	fileName := op2
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	items := make([]opts.BarData, 0)
	for i := 0; i < 6; i++ {
		for ix, row := range content {
			if ix == fila {
				for _, valor := range row {
					items = append(items, opts.BarData{Value: valor})
				}
			}
		}
	}
	return items
}

//-----------------------------------------------------------------------------//

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func archivoCSVLineItems(op2 string, fila int) []opts.LineData {
	fileName := op2
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		for ix, row := range content {
			if ix == fila {
				for _, valor := range row {
					items = append(items, opts.LineData{Value: valor})
				}
			}
		}
	}
	return items
}

//-----------------------------------------------------------------------------//

func generatePieItems() []opts.PieData {
	dias := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	items := make([]opts.PieData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.PieData{
			Name:  dias[i],
			Value: rand.Intn(300)})
	}
	return items
}

func archivoCSVPieItems(op2 string, fila int) []opts.PieData {
	fileName := op2
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not readall, err is %+v", err)
	}
	dias := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	items := make([]opts.PieData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.PieData{
			Name:  dias[i],
			Value: content[fila][i]})
	}
	return items
}

//-----------------------------------------------------------------------------//

func httpserverBar(f http.ResponseWriter, _ *http.Request) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		Theme: types.ThemeChalk,
	}),
		charts.WithTitleOpts(opts.Title{
			Title:    "GRAFICO DE BARRAS",
			Subtitle: "Con datos generados aleatoriamente",
		}))
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	os.Create("bar.html")
	bar.Render(f)
}

func httpserverBarCSV(f http.ResponseWriter, _ *http.Request) {
	data := datos{
		op2: os.Args[2],
	}
	fileName := data.op2
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	SliceLetras := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "j", "K"}
	if err != nil {
		log.Fatalf("No se pudo leer el archivo, el error es: %+v", err)
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		Theme: types.ThemeChalk,
	}),
		charts.WithTitleOpts(opts.Title{
			Title:    "GRAFICA DE BARRAS",
			Subtitle: "Con datos tomados de un archivo CSV",
		}))

	for ix, _ := range content {
		bar.SetXAxis([]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}).
			AddSeries("Category"+SliceLetras[ix], archivoCSVBarItems(data.op2, ix))
	}
	os.Create("barCSV.html")
	bar.Render(f)
}

//-----------------------------------------------------------------------------//

func httpserverLine(f http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeChalk,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "GRAFICO DE LINEAS",
			Subtitle: "Con datos generados aleatoriamente",
		}),
	)
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	os.Create("line.html")
	line.Render(f)
}

func httpserverLineCSV(f http.ResponseWriter, _ *http.Request) {
	data := datos{op2: os.Args[2]}
	fileName := data.op2
	fs1, _ := os.Open(fileName)
	r1 := csv.NewReader(fs1)
	content, err := r1.ReadAll()
	SliceLetras := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "j", "K"}
	if err != nil {
		log.Fatalf("No se pudo leer el archivo, el error es: %+v", err)
	}
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeChalk,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "GRAFICO DE LINEAS",
			Subtitle: "Con datos tomados de un archivo CSV",
		}),
	)
	for ix, _ := range content {
		line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
			AddSeries("Category "+SliceLetras[ix], archivoCSVLineItems(data.op2, ix)).
			SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	}
	os.Create("lineCSV.html")
	line.Render(f)
}

//-----------------------------------------------------------------------------//

func httpserverPie(f http.ResponseWriter, _ *http.Request) {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		Theme: types.ThemeChalk,
	}),
		charts.WithTitleOpts(
			opts.Title{
				Title:    "GRAFICA DE PASTEL",
				Subtitle: "Con datos generados aleatoriamente",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		generatePieItems()).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: 200,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				},
			),
		)
	pie.Render(f)
}

func httpserverPieCSV(f http.ResponseWriter, _ *http.Request) {
	data := datos{
		op2: os.Args[2],
	}
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeChalk,
		}),
		charts.WithTitleOpts(
			opts.Title{
				Title:    "GRAFICA DE PASTEL",
				Subtitle: "Con datos tomados de un archivo CSV",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Monthly revenue",
		archivoCSVPieItems(data.op2, 0)).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: 200,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				},
			),
		)
	pie.Render(f)
}

//-----------------------------------------------------------------------------//

var plantilla = template.Must(template.ParseGlob("*.html"))

func bar(w http.ResponseWriter, r *http.Request) {
	plantilla.ExecuteTemplate(w, "bar.html", nil)
}

func line(w http.ResponseWriter, r *http.Request) {
	plantilla.ExecuteTemplate(w, "line.html", nil)
}

func pie(w http.ResponseWriter, r *http.Request) {
	plantilla.ExecuteTemplate(w, "pie.html", nil)
}

func main() {
	if len(os.Args) == 2 {
		data1 := datos{
			op1: os.Args[1],
		}
		switch data1.op1 {
		case "--showgraphs":
			createHTMLgraphs()
			log.Println("Servidor iniciado en http://localhost:8080")
			log.Println("Grafica de barras en http://localhost:8080/bar")
			http.HandleFunc("/bar", bar)
			log.Println("Grafica de lineas en http://localhost:8080/line")
			http.HandleFunc("/line", line)
			log.Println("Grafica de pastel en http://localhost:8080/pie")
			http.HandleFunc("/pie", pie)
			http.ListenAndServe(":8080", nil)
		default:
			fmt.Println("Opcion no encontrada")
		}
	} else if len(os.Args) == 3 {
		data2 := datos{
			op1: os.Args[1],
			op2: os.Args[2],
		}
		switch data2.op1 {
		case "--bar":
			fmt.Println("Creacion de grafica de barras")
			switch data2.op2 {
			case "--generate":
				log.Println("Servidor iniciado en http://localhost:8080/bar-random")
				http.HandleFunc("/", httpserverBar)
				http.ListenAndServe(":8080", nil)
			default:
				log.Println("Servidor iniciado en http://localhost:8080/bar-csv")
				http.HandleFunc("/", httpserverBarCSV)
				http.ListenAndServe(":8080", nil)
			}
		case "--pie":
			fmt.Println("Grafcia de pastel")
			switch data2.op2 {
			case "--generate":
				log.Println("Servidor iniciado en http://localhost:8080/")
				http.HandleFunc("/", httpserverPie)
				http.ListenAndServe(":8080", nil)
			default:
				log.Println("Servidor iniciado en http://localhost:8080/")
				http.HandleFunc("/", httpserverPieCSV)
				http.ListenAndServe(":8080", nil)
			}
		case "--line":
			fmt.Println("Grafcia de lineas")
			switch data2.op2 {
			case "--generate":
				log.Println("Servidor iniciado en http://localhost:8080/")
				http.HandleFunc("/", httpserverLine)
				http.ListenAndServe(":8080", nil)
			default:
				log.Println("Servidor iniciado en http://localhost:8080/")
				http.HandleFunc("/", httpserverLineCSV)
				http.ListenAndServe(":8080", nil)
			}
		default:
			fmt.Println("Opcion no encontrada")
		}
	} else if len(os.Args) == 7 {
		filas, err := strconv.Atoi(os.Args[2])
		columnas, err2 := strconv.Atoi(os.Args[3])
		valMaximo, err3 := strconv.Atoi(os.Args[5])
		valMinimo, err4 := strconv.Atoi(os.Args[6])
		if err != nil { //Se hacen algunas validaciones por si el usuario ingresa datos erroneos
			fmt.Println(" Porfavor introduce un valor numerico, error en: '", os.Args[1], "'")
		} else if err2 != nil {
			fmt.Println(" Porfavor introduce un valor numerico, error en: '", os.Args[2], "'")
		} else if err3 != nil {
			fmt.Println(" Porfavor introduce un valor numerico, error en: '", os.Args[3], "'")
		} else if err4 != nil {
			fmt.Println(" Porfavor introduce un valor numerico, error en: '", os.Args[4], "'")
		} else {
			data := datosCSV{
				op1:       os.Args[1],
				filas:     filas,
				columnas:  columnas,
				nombre:    os.Args[4],
				valMaximo: valMaximo,
				valMinimo: valMinimo,
			}
			switch data.op1 {
			case "--createCSV":
				var miSlice = make([][]string, filas) //Se crea el slice 2d
				for i := 0; i < filas; i++ {
					miSlice[i] = make([]string, columnas)
				}
				rand.Seed(time.Now().UnixNano())
				for i := 0; i < filas; i++ {
					for j := 0; j < columnas; j++ { //Ciclo para rellener el resto de columnas
						numeros := valMinimo + rand.Intn(valMaximo-valMinimo)
						strr := strconv.Itoa(numeros) //Se convierten a string para ser ingresados al slice
						miSlice[i][j] = strr
					}
				}
				file, err := os.Create(data.nombre) //Se crea el archivo
				if err != nil {
					log.Fatal(" Ha ocurrido un error", err)
				}
				archivo := csv.NewWriter(file)
				defer archivo.Flush()
				for _, rows := range miSlice {
					archivo.Write(rows)
				}
				fmt.Println("Archivo CSV creado exitosamente")
			default:
				fmt.Println("Opcion no encontrada")
			}
		}
	}
}
