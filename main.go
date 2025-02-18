package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/sqweek/dialog"
	"github.com/xuri/excelize/v2"
)

// Expressão regular para validar nomes (somente letras, espaços e pontos)
var nomeRegex = regexp.MustCompile(`^[a-zA-ZÀ-ú\s.]+$`)

// Matérias permitidas (sem acentos e minúsculas)
var materiasValidas = map[string]bool{
	"matematica": true, "fisica": true, "quimica": true, "biologia": true,
	"redacao": true, "portugues": true, "geografia": true,
}

// Estrutura para armazenar os professores
type Professor struct {
	Nome               string `json:"nome"`
	Materia            string `json:"materia"`
	QuantidadeAulas    int    `json:"quantidade_aulas"`
	DiasDisponiveis    string `json:"dias_disponiveis"`
	HorariosDisponiveis string `json:"horarios_disponiveis"`
}

// Função para remover acentos e converter para minúsculas
func normalizarTexto(texto string) string {
	var resultado strings.Builder
	for _, char := range texto {
		switch char {
		case 'á', 'à', 'â', 'ã', 'ä', 'Á', 'À', 'Â', 'Ã', 'Ä':
			resultado.WriteRune('a')
		case 'é', 'è', 'ê', 'ë', 'É', 'È', 'Ê', 'Ë':
			resultado.WriteRune('e')
		case 'í', 'ì', 'î', 'ï', 'Í', 'Ì', 'Î', 'Ï':
			resultado.WriteRune('i')
		case 'ó', 'ò', 'ô', 'õ', 'ö', 'Ó', 'Ò', 'Ô', 'Õ', 'Ö':
			resultado.WriteRune('o')
		case 'ú', 'ù', 'û', 'ü', 'Ú', 'Ù', 'Û', 'Ü':
			resultado.WriteRune('u')
		case 'ç', 'Ç':
			resultado.WriteRune('c')
		default:
			resultado.WriteRune(unicode.ToLower(char)) // Transforma para minúsculas
		}
	}
	return resultado.String()
}

// Função para validar e converter uma linha do Excel
func linhaParaProfessor(linha []string) (*Professor, error) {
	if len(linha) < 5 {
		return nil, errors.New("linha incompleta")
	}

	// Validação do Nome: deve conter letras, espaços e pontos apenas
	nome := strings.TrimSpace(linha[0])
	if nome == "" || !nomeRegex.MatchString(nome) {
		return nil, fmt.Errorf("nome inválido: '%s'", nome)
	}

	// Validação da Matéria
	materia := strings.TrimSpace(linha[1])
	materiaNormalizada := normalizarTexto(materia) // Remover acentos e converter para minúsculas
	if !materiasValidas[materiaNormalizada] {
		return nil, fmt.Errorf("matéria inválida para '%s': '%s'", nome, materia)
	}

	// Validação da Quantidade de Aulas (deve ser um número inteiro positivo)
	quantidadeAulas, err := strconv.Atoi(strings.TrimSpace(linha[2]))
	if err != nil || quantidadeAulas <= 0 {
		return nil, fmt.Errorf("quantidade de aulas inválida para '%s': '%s'", nome, linha[2])
	}

	// Dias e Horários (não obrigatórios, mas removemos espaços extras)
	diasDisponiveis := strings.TrimSpace(linha[3])
	horariosDisponiveis := strings.TrimSpace(linha[4])

	return &Professor{
		Nome:               nome,
		Materia:            materia,
		QuantidadeAulas:    quantidadeAulas,
		DiasDisponiveis:    diasDisponiveis,
		HorariosDisponiveis: horariosDisponiveis,
	}, nil
}

func main() {
	arquivo, err := dialog.File().Filter("Arquivos Excel", "xlsx").Title("Selecione o arquivo Excel").Load()
	if err != nil {
		log.Fatal("Nenhum arquivo foi selecionado:", err)
	}

	f, err := excelize.OpenFile(arquivo)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}

	abas := f.GetSheetList()
	fmt.Println("Abas disponíveis:", abas)

	aba := "Dados"

	existe := false
	for _, nome := range abas {
		if nome == aba {
			existe = true
			break
		}
	}
	if !existe {
		log.Fatal("Aba não encontrada:", aba)
	}

	linhas, err := f.GetRows(aba)
	if err != nil {
		log.Fatal("Erro ao ler a aba:", err)
	}

	if len(linhas) < 2 {
		log.Fatal("Planilha vazia ou sem dados suficientes")
	}

	var professores []Professor
	errorsOccurred := false

	for i, linha := range linhas[1:] {
		prof, err := linhaParaProfessor(linha)
		if err != nil {
			errorsOccurred = true
			fmt.Printf("⚠️ Erro na linha %d: %v\n", i+2, err)
		} else {
			professores = append(professores, *prof)
		}
	}

	if errorsOccurred {
		fmt.Println("🚨 Erros encontrados. Corrija os dados e tente novamente.")
		return
	}

	// 🔥 Convertendo para JSON e imprimindo na tela
	jsonData, err := json.MarshalIndent(professores, "", "  ")
	if err != nil {
		log.Fatal("Erro ao converter para JSON:", err)
	}

	fmt.Println("\n📌 Dados em JSON:")
	fmt.Println(string(jsonData))
}
