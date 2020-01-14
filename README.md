# RedCoinApi
O repositório RedCoinApi contém as operações para viabilizar as transações de compra e venda de bitcoins.
<h2>Sobre o desenvolvimento do projeto</h2>
<p>A versão inicial da RedCoinApi conta com as seguintes entidades para satisfazer as operações de Compra e venda de bitcoins:<br>
    <ul>
        <li>
            ClienteApi - Representa os clientes que tem acesso aos métodos da API.
        </li>
        <li>
            Usuario - Representa os usuários que realizam a operação de compra e venda de bitcoins.
        </li>
        <li>PerfilUsuario - Representa a categoria do Usuario, se ele é vendedor, comprador ou ambos.</li>
        <li>
            Operacao - Representa as movimentações de compra e venda dos bitcoins.
        </li>
        <li>
            TipoOperacao - Representa o tipo da movimentação que o Usuario está realizando.
        </li>
    </ul>
    <img src="Documentacao/RedCoinApi-Der.png" />
</p>
<h2>Utilização do Redis</h2>
A fim de Otimizar o desempenho das operações de compra e venda de BitCoin, é utilizado o banco de dados Redis, para gravar em cache-server a cotação dos Bitcoins.<br/>
Como funciona:<br>
<ol>
    <li>A regra para a precificação dos Bitcoins, é que seja atualizada a cotação de preço de 1 em 1 hora;</li>
    <li>Dessa forma, a "primeira" consulta da cotação de preço de bitcoins é armazenada no Redis, com o prazo de expiração de 50 minutos;</li>
    <li>Enquanto o valor estiver válido no cache do Redis, as operações não realizam a consulta da API para precificação de Bitcoin.</li>
</ol>
<h2>Instalando a RedCoinApi</h2>
<p>Para o funcionamento da RedCoinApi serão necessários os seguintes requisitos:
<ul>
    <li>Docker - 19.03</li>
    <li>Docker-Compose - 19.03</li>
    <li>golang - 13.5</li>
    <li>MySql Workbench ou similar (para rodar o script de migration)</li>
</ul></p>
<p>Portas utilizadas</p>
<ul>
    <li>MySql - 1805</li>
    <li>Redis - 6372</li>
    <li>RedCoinApi - 2801</li>
</ul>
<p>Siga os seguintes passos:</p>
<ol>
    <li>
        Clone o repositório <a href="http://www.github.com/rteles86/RedCoinApi" target="_blank">RedCoinApi</a>
    </li>
    <li>
        Abra o terminal de comando de seu sistema operacional. Ex: MS-DOS (Windows) ou TERMINAL - MacOS
    </li>
    <li>
        Navegue até o diretório do repositório RedCoinApi "./Instalacao/docker-compose"<br/>
        &nbsp;&nbsp;&nbsp;- Execute o seguinte comando: <b>docker-compose up -d</b>
    </li>
    <li>
        Navegue até o diretório do repositório RedCoinApi "./API"<br/>
        &nbsp;&nbsp;&nbsp;- Execute o seguinte comando: <b>docker build -t redcoin .</b>
    </li>
    <li>
        Execute a imagem da RedCoinApi<br />
        &nbsp;&nbsp;&nbsp;- Execute o seguinte comando: <b>docker run -d -p 2801:2801 --name apiredcoin redcoin</b>
    </li>
    <li>
        Execute o arquivo de migration<br />
        &nbsp;&nbsp;&nbsp;Navegue até o diretório "./Instalacao/migration"<br />
        &nbsp;&nbsp;&nbsp;Execute o script na instância do MySql criada após a execução do "docker-compose/mysql"<br />
    </li>
</ol>
<h2>
    Exemplo de consumo dos Endpoints RedCoinApi
</h2>
<p>Após a instalação da RedCoinApi, recomendo a utilização da ferramenta Postman, para realizar os testes e consumo dos Endpoints.</p>
<h3>Adicionar Cliente Api</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Realiza o cadastro de um novo cliente da RedCoinApi<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/autenticacao/adicionar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: POST<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"usuario":"redcoinApi"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;,"senha":"redcoinApi@@"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
<h3>Autenticar Cliente Api</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Verifica se as credenciais do cliente da RedCoinApi são válidas<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/autenticacao<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: POST<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"usuario":"redcoinApi"
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;,"senha":"redcoinApi@@"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
<h3>Listar Todos Perfil</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Lista de todos os Perfil cadastrados<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/perfil/todos<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET
</p>
<h3>Perfil Id</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Retorna o Perfil cadastrado de acordo com o ID<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/perfil/id<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros: ?id=1
</p>
<h3>Perfil Adicionar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Realiza o cadastro de um novo Perfil<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/perfil/adicionar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: POST<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"perfil":"Comprador"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
<h3>Perfil Alterar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Altera o registro de um Perfil<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/perfil/alterar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: PUT<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"id":1
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;,"perfil":"Corretor"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>


<h3>Listar Todos TipoOperacao</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Lista de todos os TipoOperacao cadastrados<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/tipo-operacao/todos<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET
</p>
<h3>TipoOperacao Id</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Retorna o TipoOperacao cadastrado de acordo com o ID<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/tipo-operacao/id<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros: ?id=1
</p>
<h3>TipoOperacao Adicionar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Realiza o cadastro de um novo TipoOperacao<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/tipo-operacao/adicionar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: POST<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"operacao":"Venda"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
<h3>TipoOperacao Alterar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Altera o registro de um TipoOperacao<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/tipo-operacao/alterar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: PUT<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"id":1
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;,"operacao":"Compra"
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>


<h3>Listar Todos Usuario</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Lista de todos os Usuario cadastrados<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/usuario/todos<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET
</p>
<h3>Usuario Id</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Retorna o Usuario cadastrado de acordo com o ID<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/usuario/id?id={valor}<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros: ?id=1
</p>
<h3>Usuario Adicionar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Realiza o cadastro de um novo Usuario<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/usuario/adicionar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: POST<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"email": "redcoinapi@redcoinapi.com",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"senha": "123Mudar",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"nome": "Red Coin",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"ultimoNome": "Api",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"nascimento": "2019-12-22T00:00:00Z",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"quantidadeMoeda": 0,
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"perfil": [
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"id": 1,
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"perfil": "Comprador"
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;]
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
<h3>Usuario Alterar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Altera o registro de um Usuario<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/usuario/alterar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: PUT<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros Body: 
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"id": 1,
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"email": "redcoinapi@redcoinapi.com",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"senha": "123Mudar",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"nome": "Red Coin",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"ultimoNome": "Api",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"nascimento": "2019-12-22T00:00:00Z",
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"quantidadeMoeda": 0,
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"perfil": [
        <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"id": 2,
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"perfil": "Vendedor"
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
            <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;]
    <br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
<h3>Operacao Email</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Lista todas as transações de um Usuario a partir de seu Email<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/operacao/email?email={valor}<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros: ?email=redcoinapi@redcoinapi.com
</p>
<h3>Operacao Data</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Lista todas as transações registradas em uma determinada data<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/operacao/data?data={valor}<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: GET<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros: ?data=2019-12-22
</p>
<h3>Operacao Gravar</h3>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Descrição: Retorna o Perfil cadastrado de acordo com o ID<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;URL: http://localhost:2801/api/operacao/gravar<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Tipo: POST<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Parâmetros body:<br>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{<br>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"idTipo": 1,<br>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"idVendedor": 2,<br>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"idComprador": 2{<br>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"DataOperacao": "2019-12-22T01:56:32Z",<br>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"ValorMoeda": 14717.366,<br>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"ValorBitCoin": 0.10<br>
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;}
</p>
