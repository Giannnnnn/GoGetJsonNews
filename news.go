package main

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Title       string `xml:"title"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

func fetchRSSFeed(url string) (RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()

	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return RSS{}, err
	}

	return rss, nil
}

func topicHandler(c *gin.Context) {
	platform := c.Param("platform")
	topic := c.Param("topic")

	var url string
	switch platform {
	case "g1":
		url = g1Topics[topic]
	case "regions":
		url = g1Regions[topic]
	default:
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Platform not found",
		})
		return
	}

	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Topic not found",
		})
		return
	}

	rss, err := fetchRSSFeed(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch RSS feed",
		})
		return
	}

	c.JSON(http.StatusOK, rss.Channel.Items)
}

var g1Topics = map[string]string{
	"brasil":             "https://g1.globo.com/rss/g1/brasil/",
	"carros":             "https://g1.globo.com/rss/g1/carros/",
	"ciencia-e-saude":    "https://g1.globo.com/rss/g1/ciencia-e-saude/",
	"concursos-e-emprego": "https://g1.globo.com/rss/g1/concursos-e-emprego/",
	"economia":           "https://g1.globo.com/dynamo/economia/rss2.xml",
	"educacao":           "https://g1.globo.com/dynamo/educacao/rss2.xml",
	"loterias":           "https://g1.globo.com/dynamo/loterias/rss2.xml",
	"mundo":              "https://g1.globo.com/dynamo/mundo/rss2.xml",
	"musica":             "https://g1.globo.com/dynamo/musica/rss2.xml",
	"natureza":           "https://g1.globo.com/dynamo/natureza/rss2.xml",
	"planeta-bizarro":    "https://g1.globo.com/dynamo/planeta-bizarro/rss2.xml",
	"politica":           "https://g1.globo.com/dynamo/politica/mensalao/rss2.xml",
	"pop-arte":           "https://g1.globo.com/dynamo/pop-arte/rss2.xml",
	"tecnologia":         "https://g1.globo.com/dynamo/tecnologia/rss2.xml",
	"turismo-e-viagem":   "https://g1.globo.com/dynamo/turismo-e-viagem/rss2.xml",
}
var g1Regions = map[string]string{
	"acre":                        "https://g1.globo.com/dynamo/ac/acre/rss2.xml",
	"alagoas":                     "https://g1.globo.com/dynamo/al/alagoas/rss2.xml",
	"amapa":                       "https://g1.globo.com/dynamo/ap/amapa/rss2.xml",
	"amazonas":                    "https://g1.globo.com/dynamo/am/amazonas/rss2.xml",
	"bahia":                       "https://g1.globo.com/dynamo/bahia/rss2.xml",
	"ceara":                       "https://g1.globo.com/dynamo/ceara/rss2.xml",
	"distrito-federal":            "https://g1.globo.com/dynamo/distrito-federal/rss2.xml",
	"espirito-santo":              "https://g1.globo.com/dynamo/espirito-santo/rss2.xml",
	"goias":                       "https://g1.globo.com/dynamo/goias/rss2.xml",
	"maranhao":                    "https://g1.globo.com/dynamo/ma/maranhao/rss2.xml",
	"mato-grosso":                 "https://g1.globo.com/dynamo/mato-grosso/rss2.xml",
	"mato-grosso-do-sul":          "https://g1.globo.com/dynamo/mato-grosso-do-sul/rss2.xml",
	"minas-gerais":                "https://g1.globo.com/dynamo/minas-gerais/rss2.xml",
	"mg-centro-oeste":             "https://g1.globo.com/dynamo/mg/centro-oeste/rss2.xml",
	"mg-grande-minas":             "https://g1.globo.com/dynamo/mg/grande-minas/rss2.xml",
	"mg-sul-de-minas":             "https://g1.globo.com/dynamo/mg/sul-de-minas/rss2.xml",
	"mg-triangulo-mineiro":        "https://g1.globo.com/dynamo/minas-gerais/triangulo-mineiro/rss2.xml",
	"mg-vales-de-minas-gerais":    "https://g1.globo.com/dynamo/mg/vales-mg/rss2.xml",
	"mg-zona-da-mata":             "https://g1.globo.com/dynamo/mg/zona-da-mata/rss2.xml",
	"para":                        "https://g1.globo.com/dynamo/pa/para/rss2.xml",
	"paraiba":                     "https://g1.globo.com/dynamo/pb/paraiba/rss2.xml",
	"parana":                      "https://g1.globo.com/dynamo/pr/parana/rss2.xml",
	"pr-campos-gerais-sul":        "https://g1.globo.com/dynamo/pr/campos-gerais-sul/rss2.xml",
	"pr-oeste-sudoeste":           "https://g1.globo.com/dynamo/pr/oeste-sudoeste/rss2.xml",
	"pr-norte-noroeste":           "https://g1.globo.com/dynamo/pr/norte-noroeste/rss2.xml",
	"pernambuco":                  "https://g1.globo.com/dynamo/pernambuco/rss2.xml",
	"pe-caruaru-regiao":           "https://g1.globo.com/dynamo/pe/caruaru-regiao/rss2.xml",
	"pe-petrolina-regiao":         "https://g1.globo.com/dynamo/pe/petrolina-regiao/rss2.xml",
	"rio-de-janeiro":              "https://g1.globo.com/dynamo/rio-de-janeiro/rss2.xml",
	"rj-regiao-serrana":           "https://g1.globo.com/dynamo/rj/regiao-serrana/rss2.xml",
	"rj-regiao-dos-lagos":         "https://g1.globo.com/dynamo/rj/regiao-dos-lagos/rss2.xml",
	"rj-norte-fluminense":         "https://g1.globo.com/dynamo/rj/norte-fluminense/rss2.xml",
	"rj-sul-do-rio-costa-verde":   "https://g1.globo.com/dynamo/rj/sul-do-rio-costa-verde/rss2.xml",
	"rio-grande-do-norte":         "https://g1.globo.com/dynamo/rn/rio-grande-do-norte/rss2.xml",
	"rio-grande-do-sul":           "https://g1.globo.com/dynamo/rs/rio-grande-do-sul/rss2.xml",
	"rondonia":                    "https://g1.globo.com/dynamo/ro/rondonia/rss2.xml",
	"roraima":                     "https://g1.globo.com/dynamo/rr/roraima/rss2.xml",
	"santa-catarina":              "https://g1.globo.com/dynamo/sc/santa-catarina/rss2.xml",
	"sao-paulo":                   "https://g1.globo.com/dynamo/sao-paulo/rss2.xml",
	"sp-bauru-marilia":            "https://g1.globo.com/dynamo/sp/bauru-marilia/rss2.xml",
	"sp-campinas-regiao":          "https://g1.globo.com/dynamo/sp/campinas-regiao/rss2.xml",
	"sp-itapetininga-regiao":      "https://g1.globo.com/dynamo/sao-paulo/itapetininga-regiao/rss2.xml",
	"sp-mogi-das-cruzes-suzano":   "https://g1.globo.com/dynamo/sp/mogi-das-cruzes-suzano/rss2.xml",
	"sp-piracicaba-regiao":        "https://g1.globo.com/dynamo/sp/piracicaba-regiao/rss2.xml",
	"sp-presidente-prudente-regiao": "https://g1.globo.com/dynamo/sp/presidente-prudente-regiao/rss2.xml",
	"sp-ribeirao-preto-franca":    "https://g1.globo.com/dynamo/sp/ribeirao-preto-franca/rss2.xml",
	"sp-sao-jose-do-rio-preto-aracatuba": "https://g1.globo.com/dynamo/sao-paulo/sao-jose-do-rio-preto-aracatuba/rss2.xml",
	"sp-santos-regiao":            "https://g1.globo.com/dynamo/sp/santos-regiao/rss2.xml",
	"sp-sao-carlos-regiao":        "https://g1.globo.com/dynamo/sp/sao-carlos-regiao/rss2.xml",
	"sp-sorocaba-jundiai":         "https://g1.globo.com/dynamo/sao-paulo/sorocaba-jundiai/rss2.xml",
	"sp-vale-do-paraiba-regiao":   "https://g1.globo.com/dynamo/sp/vale-do-paraiba-regiao/rss2.xml",
	"sergipe":                     "https://g1.globo.com/dynamo/se/sergipe/rss2.xml",
	"tocantins":                   "https://g1.globo.com/dynamo/to/tocantins/rss2.xml",
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/topics/:platform/:topic", topicHandler)

	router.Run(":8080")
}
