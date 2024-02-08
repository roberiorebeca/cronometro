package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title                string
	ButtonsDiscurso      []Button
	ButtonsAparte        []Button
	ButtonsOrdem         []Button
	ButtonsConsideracoes []Button
	Message              string // Mensagem para exibir após a ação
}

type Button struct {
	Label string
	URL   string
	Cor   string
}

var templates *template.Template

func main() {
	http.HandleFunc("/", HomePage)

	http.HandleFunc("/discursoIniciar", ButtonDiscursoIniciar)
	http.HandleFunc("/discursoPararZerar", ButtonDiscursoPararZerar)
	http.HandleFunc("/discursoAdicionarMinuto", ButtonDiscursoAdicionarMinuto)
	http.HandleFunc("/discursoParar", ButtonDiscursoParar)

	http.HandleFunc("/aparteIniciar", ButtonAparteIniciar)
	http.HandleFunc("/apartePararZerar", ButtonApartePararZerar)
	http.HandleFunc("/aparteParar", ButtonAparteParar)

	http.HandleFunc("/ordemIniciar", ButtonOrdemIniciar)
	http.HandleFunc("/ordemPararZerar", ButtonOrdemPararZerar)
	http.HandleFunc("/ordemParar", ButtonOrdemParar)

	http.HandleFunc("/consideracoesIniciar", ButtonConsideracoesIniciar)
	http.HandleFunc("/consideracoesPararZerar", ButtonConsideracoesPararZerar)
	http.HandleFunc("/consideracoesParar", ButtonConsideracoesParar)

	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	pageVariables := PageVariables{
		Title: "Cronômetro",
		ButtonsDiscurso: []Button{
			{Label: "Iniciar", URL: "/discursoIniciar", Cor: "success"},
			{Label: "Parar/Zerar", URL: "/discursoPararZerar", Cor: "warning"},
			{Label: "+1 Minuto", URL: "/discursoAdicionarMinuto", Cor: "primary"},
			{Label: "Parar", URL: "/discursoParar", Cor: "danger"},
		},
		ButtonsAparte: []Button{
			{Label: "Iniciar", URL: "/aparteIniciar", Cor: "success"},
			{Label: "Parar/Zerar", URL: "/apartePararZerar", Cor: "warning"},
			{Label: "Parar", URL: "/aparteParar", Cor: "danger"},
		},
		ButtonsOrdem: []Button{
			{Label: "Iniciar", URL: "/ordemIniciar", Cor: "success"},
			{Label: "Parar/Zerar", URL: "/ordemPararZerar", Cor: "warning"},
			{Label: "Parar", URL: "/ordemParar", Cor: "danger"},
		},
		ButtonsConsideracoes: []Button{
			{Label: "Iniciar", URL: "/consideracoesIniciar", Cor: "success"},
			{Label: "Parar/Zerar", URL: "/consideracoesPararZerar", Cor: "warning"},
			{Label: "Parar", URL: "/consideracoesParar", Cor: "danger"},
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

// Comandos do Discurso
func ButtonDiscursoIniciar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fc=CONTAR", w)

	http.Redirect(w, r, "/?message=Iniciado com sucesso", http.StatusFound)

}

func ButtonDiscursoPararZerar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Parado e Zerado com sucesso", http.StatusFound)
}

func ButtonDiscursoAdicionarMinuto(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=1&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=1&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/discursoIniciar", http.StatusFound)
}

func ButtonDiscursoParar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=8&n6=0&Fp=PARAR", w)

	http.Redirect(w, r, "/?message=Parado com sucesso", http.StatusFound)
}

// Comandos do Aparte
func ButtonAparteIniciar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fc=CONTAR", w)

	http.Redirect(w, r, "/?message=Iniciado com sucesso", http.StatusFound)

}

func ButtonApartePararZerar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Parado e Zerado com sucesso", http.StatusFound)
}

func ButtonAparteParar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fp=PARAR", w)

	http.Redirect(w, r, "/?message=Parado com sucesso", http.StatusFound)
}

// Comandos da Questão de Ordem
func ButtonOrdemIniciar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fc=CONTAR", w)

	http.Redirect(w, r, "/?message=Iniciado com sucesso", http.StatusFound)

}

func ButtonOrdemPararZerar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Parado e Zerado com sucesso", http.StatusFound)
}

func ButtonOrdemParar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=2&n6=0&Fp=PARAR", w)

	http.Redirect(w, r, "/?message=Parado com sucesso", http.StatusFound)
}

// Comandos da Considerações Finais
func ButtonConsideracoesIniciar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=3&n6=0&Fc=CONTAR", w)

	http.Redirect(w, r, "/?message=Iniciado com sucesso", http.StatusFound)

}

func ButtonConsideracoesPararZerar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=3&n6=0&Fp=PARAR", w)
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=3&n6=0&Fz=ZERAR", w)

	http.Redirect(w, r, "/?message=Parado e Zerado com sucesso", http.StatusFound)
}

func ButtonConsideracoesParar(w http.ResponseWriter, r *http.Request) {
	executarURL("http://192.168.1.80/Bcron.htm?C1=2&f2=on&n4=0&n5=3&n6=0&Fp=PARAR", w)

	http.Redirect(w, r, "/?message=Parado com sucesso", http.StatusFound)
}
