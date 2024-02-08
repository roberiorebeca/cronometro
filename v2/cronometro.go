package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title             string
	ButtonsConfig     []Button
	ButtonsComando    []Button
	ButtonsAdicionais []Button
	Message           string // Mensagem para exibir após a ação
}

type Button struct {
	Label string
	URL   string
	Cor   string
}

var templates *template.Template

func main() {
	http.HandleFunc("/", HomePage)

	http.HandleFunc("/discursoConfig", ButtonDiscurso)
	http.HandleFunc("/aparteConfig", ButtonAparte)
	http.HandleFunc("/ordemConfig", ButtonOrdem)
	http.HandleFunc("/consideracoesConfig", ButtonConsideracoes)

	http.HandleFunc("/comandoIniciar", ButtonIniciar)
	http.HandleFunc("/comandoParar", ButtonParar)

	http.HandleFunc("/comandoAdicionarMinuto", ButtonAdicionarMinuto)

	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	pageVariables := PageVariables{
		Title: "Cronômetro",
		ButtonsConfig: []Button{
			{Label: "Discurso", URL: "/discursoConfig", Cor: "success"},
			{Label: "Aparte", URL: "/aparteConfig", Cor: "warning"},
			{Label: "Questão de Ordem", URL: "/ordemConfig", Cor: "primary"},
			{Label: "Considerações Finais", URL: "/consideracoesConfig", Cor: "danger"},
		},
		ButtonsComando: []Button{
			{Label: "Iniciar", URL: "/comandoIniciar", Cor: "success"},
			{Label: "Parar", URL: "/comandoParar", Cor: "danger"},
		},
		ButtonsAdicionais: []Button{
			{Label: "+1 Minuto", URL: "/comandoAdicionarMinuto", Cor: "info"},
		},
		Message: message,
	}

	templates = template.Must(template.ParseGlob("*.html"))

	templates.ExecuteTemplate(w, "home.html", pageVariables)
}

func executarURL(url string, w http.ResponseWriter) {
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao acessar a URL: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
}

// Comandos da Configuração do Tipo de Cronometro
func ButtonDiscurso(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Configurado para Discurso", http.StatusFound)
}

func ButtonAparte(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Configurado para Aparte", http.StatusFound)
}

func ButtonOrdem(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Configurado para Ordem", http.StatusFound)
}

func ButtonConsideracoes(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=3&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=3&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Configurado para as Considerações", http.StatusFound)
}

// Comandos do Cronometro
func ButtonIniciar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fc=CONTAR", w)

	http.Redirect(w, r, "/?message=Iniciado com sucesso", http.StatusFound)

}

func ButtonParar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fp=PARAR", w)

	http.Redirect(w, r, "/?message=Parado com sucesso", http.StatusFound)
}

func ButtonAdicionarMinuto(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=1&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=1&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Adicionado 1 Minuto", http.StatusFound)
}
