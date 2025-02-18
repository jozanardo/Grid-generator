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
var nameRegex = regexp.MustCompile(`^[a-zA-ZÀ-ú\s.]+$`)

// Matérias permitidas (sem acentos e minúsculas)
var validSubjects = map[string]bool{
	"matematica": true, "fisica": true, "quimica": true, "biologia": true,
	"redacao": true, "portugues": true, "geografia": true,
}

// Estrutura para armazenar os professoes
type Teacher struct {
	Name               string 
	Subject            string 
	NumberOfClasses    int    
	AvailableDays      string 
	AvailableHours     string 
}

// Função para remover acentos e converter para minúsculas
func normalizeText(texto string) string {
	var result strings.Builder
	for _, char := range texto {
		switch char {
		case 'á', 'à', 'â', 'ã', 'ä', 'Á', 'À', 'Â', 'Ã', 'Ä':
			result.WriteRune('a')
		case 'é', 'è', 'ê', 'ë', 'É', 'È', 'Ê', 'Ë':
			result.WriteRune('e')
		case 'í', 'ì', 'î', 'ï', 'Í', 'Ì', 'Î', 'Ï':
			result.WriteRune('i')
		case 'ó', 'ò', 'ô', 'õ', 'ö', 'Ó', 'Ò', 'Ô', 'Õ', 'Ö':
			result.WriteRune('o')
		case 'ú', 'ù', 'û', 'ü', 'Ú', 'Ù', 'Û', 'Ü':
			result.WriteRune('u')
		case 'ç', 'Ç':
			result.WriteRune('c')
		default:
			result.WriteRune(unicode.ToLower(char)) // Transforma para minúsculas
		}
	}
	return result.String()
}

// Função para validar e converter uma linha do Excel
func mapLineToTeacher(line []string) (*Teacher, error) {
	if len(line) < 5 {
		return nil, errors.New("Linha incompleta")
	}

	// Validação do Nome: deve conter letras, espaços e pontos apenas
	name := strings.TrimSpace(line[0])
	if name == "" || !nameRegex.MatchString(name) {
		return nil, fmt.Errorf("Nome inválido: '%s'", name)
	}

	// Validação da Matéria
	subject := strings.TrimSpace(line[1])
	normalizedSubject := normalizeText(subject) // Remover acentos e converter para minúsculas
	if !validSubjects[normalizedSubject] {
		return nil, fmt.Errorf("matéria inválida para '%s': '%s'", name, subject)
	}

	// Validação da Quantidade de Aulas (deve ser um número inteiro positivo)
	numberOfClasses, err := strconv.Atoi(strings.TrimSpace(line[2]))
	if err != nil || numberOfClasses <= 0 {
		return nil, fmt.Errorf("quantidade de aulas inválida para '%s': '%s'", name, line[2])
	}

	// Dias e Horários (não obrigatórios, mas removemos espaços extras)
	availableDays := strings.TrimSpace(line[3])
	availableHours := strings.TrimSpace(line[4])

	return &Teacher{
		Name:               name,
		Subject:            subject,
		NumberOfClasses:    numberOfClasses,
		AvailableDays:      availableDays,
		AvailableHours: 		availableHours,
	}, nil
}

func main() {
	file, err := dialog.File().Filter("Arquivos Excel", "xlsx").Title("Selecione o arquivo Excel").Load()
	if err != nil {
		log.Fatal("Nenhum arquivo foi selecionado:", err)
	}

	f, err := excelize.OpenFile(file)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}

	tabs := f.GetSheetList()
	fmt.Println("Abas disponíveis:", tabs)

	tab := "Dados"

	exists := false
	for _, name := range tabs {
		if name == tab {
			exists = true
			break
		}
	}
	if !exists {
		log.Fatal("Aba não encontrada:", tab)
	}

	lines, err := f.GetRows(tab)
	if err != nil {
		log.Fatal("Erro ao ler a aba:", err)
	}

	if len(lines) < 2 {
		log.Fatal("Planilha vazia ou sem dados suficientes")
	}

	var Teachers []Teacher
	errorsOccurred := false

	for i, line := range lines[1:] {
		teacher, err := mapLineToTeacher(line)
		if err != nil {
			errorsOccurred = true
			fmt.Printf("⚠️ Erro na linha %d: %v\n", i+2, err)
		} else {
			Teachers = append(Teachers, *teacher)
		}
	}

	if errorsOccurred {
		fmt.Println("🚨 Erros encontrados. Corrija os dados e tente novamente.")
		return
	}

	// 🔥 Convertendo para JSON e imprimindo na tela
	jsonData, err := json.MarshalIndent(Teachers, "", "  ")
	if err != nil {
		log.Fatal("Erro ao converter para JSON:", err)
	}

	fmt.Println("\n📌 Dados em JSON:")
	fmt.Println(string(jsonData))
}
