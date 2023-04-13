# Entradas e saídas em programas em Go

## Sobre

Este conjunto de programas serve à demonstração das diferentes formas de prover entradas e saídas utilizando a linguagem de programação [Go](https://go.dev). Esses exemplos cobrem essencialmente três formas:

1. Entrada e saída padrão (*console*)
2. Argumentos de linha de comando
3. Arquivos

## Estrutura do repositório

Este repositório contém um conjunto de arquivos fonte demonstrando, cada um, o uso de diferentes pacotes, funções e métodos para possibilitar entrada e saída de dados em programas em Go. Os arquivos estão organizados de acordo com a seguinte estrutura:

```
+─io                    ---> Nome do diretório do projeto
  ├─── cli              ---> Diretório com exemplos sobre entrada utilizando linha de comando
       ├─── args        ---> Diretório com exemplo utilizando argumentos de linha de comando
       ├─── flags       ---> Diretório com exemplo utilizando flags na linha de comando
  ├─── file             ---> Diretório com exemplos sobre entrada e saída utilizando arquivos
       ├─── bufio       ---> Diretório com exemplo de leitura de arquivo com bufio.Reader.ReadString
       ├─── fileout     ---> Diretório com exemplo de escrita em arquivo com descritor
       ├─── ioutil      ---> Diretório com exemplo de leitura de arquivo com io/ioutil.ReadFile
       ├─── reader      ---> Diretório com exemplo de leitura de um arquivo grande com bufio.Reader.Read
       ├─── writer      ---> Diretório com exemplo de escrita em arquivo com bufio.Writer.Write
       ├─── writestring ---> Diretório com exemplo de escrita em arquivo com descritor
  ├─── stdin            ---> Diretório com exemplos sobre entrada via console
       ├─── readstring  ---> Diretório com exemplo de leitura da entrada padrão com bufio.Reader.ReadString
       ├─── scanln      ---> Diretório com exemplo de leitura da entrada padrão com fmt.Scanln
       ├─── scanner     ---> Diretório com exemplo de leitura da entrada padrão com fmt.Sscanf
       ├─── sscanf      ---> Diretório com exemplo de leitura da entrada padrão com bufio.Scanner
  ├─── stdout           ---> Diretório com exemplos sobre saída via console
       ├─── writer      ---> Diretório com exemplo de escrita na saída padrão com bufio.Writer.WriteString
```

## Requisitos

Para a compilação e execução dos programas, os seguintes elementos devem estar devidamente instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

## Download, compilação e execução dos programas

No terminal do sistema operacional, insira os seguintes comandos para realizar o *download* da implementação a partir deste repositório Git:

```bash
# Download da implementação a partir do repositório Git
git clone https://github.com/ufrn-golang/io.git
```

Para executar um programa, deve-se primeiro navegar para o respectivo diretório no qual ele se encontra e utilizar o comando `go run` no terminal do sistema operacional. Por exemplo, para executar o programa referente ao arquivo [`stdin/scanln`](https://github.com/ufrn-golang/io/tree/master/stdin/scanln/scanln.go), deve-se utilizar os seguintes comandos:

```bash
# Navegar para o diretório do programa
cd stdin/scanln

# Compilar e executar o programa
go run scanln.go
```

Caso deseje gerar um executável para o programa em questão, deve-se utilizar o comando `go build` e, em seguida, invocar o arquivo executável gerado. Por exemplo, Por exemplo, para compilar e executar o programa referente ao arquivo [`stdin/scanln`](https://github.com/ufrn-golang/io/tree/master/stdin/scanln/scanln.go), deve-se utilizar os seguintes comandos:

```bash
# Navegar para o diretório do programa
cd stdin/scanln

# Compilação do programa e geração de arquivo executável (com o mesmo nome do arquivo fonte)
go build scanln.go

# Execução do programa
./scanln
```
