# xy-inc
<b> Desafio LuizaLabs (xy-inc) </b>

A aplicação possui em 3 serviços: <br>
- Serviço para cadastrar os POIs <br>
- Serviço para listar todos os POIs <br>
- Serviço para buscar os POIs por proximidade, passando 3 parâmetros: <br>
> x: valor da coordenada "x" <br>
> y: valor da coordenada "y" <br>
> d: distância máxima em metros "m" <br>

<b> Iniciando a Aplicação: </b><br>
Método utilizado: POST <br>
URI: http://localhost:3000/poi <br>
Utilizando o serviço de cadastro de recurso, os seguintes POIs devem ser armazenados: 
    [
        {
            "nome": "Lanchonete",
            "coordenadaX": 27,
            "coordenadaY": 12
        },
        {
            "nome": "Posto",
            "coordenadaX": 31,
            "coordenadaY": 18
        },
        {
            "nome": "Joalheria",
            "coordenadaX": 15,
            "coordenadaY": 12
        },
        {
            "nome": "Floricultura",
            "coordenadaX": 19,
            "coordenadaY": 21
        },
        {
            "nome": "Pub",
            "coordenadaX": 12,
            "coordenadaY": 8
        },
        {
            "nome": "Supermercado",
            "coordenadaX": 23,
            "coordenadaY": 6
        },
        {
            "nome": "Churrascaria",
            "coordenadaX": 28,
            "coordenadaY": 2
        }
    ]
	
<b> Listando todos os recursos cadastrados: </b><br>
Método utilizado: GET <br>
URI: http://localhost:3000/poi

<b> Listando os recursos por proximidade </b><br>
Método utilizado: GET <br>
URI: http://localhost:3000/poi/proximidade?x=20&y=10&d=10

<b> Install </b><br>
$ go build <br>
$ ./xy-inc
