/*     // Defina a var iável "blocoId" 
    let blocoId;
 */
    let blocosSalvos = obterIdsDaUrl();
    function obterIdsDaUrl() {
        const urlParams = new URLSearchParams(window.location.search);
        const idsParam = urlParams.get('ids');

        if(idsParam) {
            return idsParam.split(',').map(id => parseInt(id));
                } else {
                    return [];
                } 
    } 



// Aguarda o carregamento completo do DOM antes de executar o script
document.addEventListener('DOMContentLoaded', function () {


// Função para obter blocos por IDs por url
function obterBlocosPorID() {
    // Obtém os dados codificados da URL da página
    const urlParams = new URLSearchParams(window.location.search);
    const ids = urlParams.get('ids');
    console.log(ids);

    // Verifica se há IDs na URL
    if (!ids) {
    console.error("IDs de blocos não encontrados na URL.");
    mostrarErroSemBlocos();
    return;
}
    // Faz uma solicitação GET para buscar os blocos correspondentes
    fetch('/buscaid?ids=' + ids, {
        method: 'GET',
        headers: {
            'X-Requested-With': 'XMLHttpRequest',
        },
    })
    .then(response => {
        // Verifica se a resposta está OK
        if (!response.ok) {
            throw new Error(`Erro na solicitação: ${response.status}`);
        }
        // Converte a resposta para JSON
        return response.json();
    })
        .then(blocos => {
            // Chama a função para adicionar os cards ao container com os blocos recebidos
            console.log("Blocos recebidos:", blocos);
            adicionarCardsAoContainer(blocos);
            adicionarOuvintesDeEventos();
        })
        .catch(error => {
            // Lida com erros durante a solicitação
            console.error("Erro durante a solicitação:", error.message);
        });
}
        
    function mostrarErroSemBlocos() {
        const blocosContainer = document.getElementById('blocosContainer');

        // Limpar qualquer conteúdo existente no contêiner
        blocosContainer.innerHTML = '';

        const tituloErro = document.createElement('h5');
        tituloErro.textContent = 'Você ainda não adicionou blocos à sua lista. Que tal buscar alguns?';
        blocosContainer.appendChild(tituloErro);

        const instrucaoErro = document.createElement('h6');
        instrucaoErro.innerHTML = 'Pra adicionar algum bloco, é só clicar no botão "<b>Salvar</b>” dentro do card do bloco escolhido. <br> Depois de salvar <b>todos</b> os blocos, clique na aba Minha Lista. 😉';
        blocosContainer.appendChild(instrucaoErro);

        const botaoBusca = document.createElement('button');
        botaoBusca.className = 'btn-blocos';
        botaoBusca.textContent = 'Buscar blocos';
        blocosContainer.appendChild(botaoBusca);

        // Adiciona um evento de clique para redirecionar para "/index.html"
        botaoBusca.addEventListener('click', function () {
        window.location.href = '/index.html';
    });
    }

    // Função para adicionar os cards ao contêiner com base nos dados recebidos
    function adicionarCardsAoContainer(blocosPorId) {
        const blocosContainer = document.getElementById('blocosContainer');

        // Limpar qualquer conteúdo existente no contêiner
        blocosContainer.innerHTML = '';

        // Adicionar os cards ao contêiner
        blocosPorId.forEach(bloco => {
            const card = criarCard(bloco);
            blocosContainer.appendChild(card);
        });
    }

    // Função para criar um card com base nos dados do bloco
    function criarCard(bloco) {
        // Cria o elemento card
        const card = document.createElement('div');
        card.className = 'card mb-2';

        // Adiciona a div com o título
        const cardTitulo = document.createElement('div');
        cardTitulo.className = 'card-titulo';
        cardTitulo.style.display = 'flex';
        cardTitulo.style.justifyContent = 'space-between';
        cardTitulo.innerHTML = `
        <h3>${bloco.Nome}</h3>
        <button class="remover btn-card" data-bloco-id="${bloco.ID}">Remover <img src="/static/images/trash.svg" alt="Excluir Icon" width="24" height="24"</img></button>
    `;
        card.appendChild(cardTitulo);

        // Adiciona a div com as informações do card
        const infoCard = document.createElement('div');
        infoCard.className = 'info-card';

        // Adiciona a div com a data
        const cardData = document.createElement('div');
        cardData.className = 'card-data col-6';
        // Obtém apenas o dia e o mês da propriedade Data
        const diaEMes = bloco.Data.substring(0, 5);

        cardData.innerHTML = `
        <img src="/static/images/calendar.svg" alt="Icon Calendario" width="24" height="24"
    </img>
            <span>${diaEMes}</span>
        `;
        infoCard.appendChild(cardData);

        // Adiciona a div com o horário
        const cardHorario = document.createElement('div');
        cardHorario.className = 'card-horario col-6';
        // Somente os 2 primeiros dígitos do horário
        const duasPrimeirasHoras = bloco.Concentracao.substring(0, 2);
        cardHorario.innerHTML = `
        <img src="/static/images/clock.svg" alt="Icon Relogio" width="24" height="24"</img>
        <span>${duasPrimeirasHoras}h</span>
        `;
        infoCard.appendChild(cardHorario);

                // Adiciona a div com o endereço
        const cardEndereco = document.createElement('div');
        cardEndereco.className = 'card-endereco endereco-column';
        cardEndereco.innerHTML = `
        <img src="/static/images/map_pin.svg" alt="Map Pin Icon" width="24" height="24"
    </img>
    <div class="endereco-content">
    <a href="https://www.google.com/maps/search/${bloco.Local}" target="_blank">${bloco.Local}</a>
    <p>${bloco.Subprefeitura}</p>
</div>
`;



        // Adiciona a div infoCard ao card
        card.appendChild(infoCard);
        card.appendChild(cardEndereco);
        return card;
    }

    obterBlocosPorID();

    function adicionarOuvintesDeEventos() {
        var buttons = document.querySelectorAll('.remover');
        buttons.forEach(button => {
            button.addEventListener('click', () => {
                let blocoId = button.getAttribute('data-bloco-id');
                removerBlocoLista(blocoId);
                console.log(blocoId);
                window.location.href = `minhalista?ids=${blocosSalvos.join(',')}`;
            });
        });
    }

   
    function removerBlocoLista(blocoId) {
        // Use filter para criar uma nova lista sem o bloco removido
        blocosSalvos = blocosSalvos.filter(id => id !== Number(blocoId));
        console.log("Blocos Salvos após a remoção:", blocosSalvos);
    }

});


function copyToClipboard() {
    try {
        // Obtém a URL da página atual
        var currentUrl = window.location.href;

        // Verifica se a URL está vazia
        if (!currentUrl) {
            throw new Error('Não foi possível obter a URL da página.');
        }

        // Define o atributo data-url no botão com a URL
        var copyButton = document.getElementById('copyButton');
        copyButton.setAttribute('data-url', currentUrl);

        // Cria um elemento input
        var dummy = document.createElement('input');
        
        // Define o valor do input para a URL atual
        dummy.value = currentUrl;

        // Adiciona o input ao documento
        document.body.appendChild(dummy);

        // Seleciona o texto do input
        dummy.select();

        // Cria um intervalo (range) para selecionar o texto
        var range = document.createRange();
        range.selectNode(dummy);

        // Cria uma seleção e a adiciona ao intervalo
        var selection = window.getSelection();
        selection.removeAllRanges();
        selection.addRange(range);

        // Copia o texto
        document.execCommand('copy');

        // Remove o input do documento
        document.body.removeChild(dummy);

        // Limpa a seleção
        selection.removeAllRanges();

        // Altera o texto do botão e desabilita
        copyButton.innerHTML = 'Link copiado <img src="/static/images/check.svg"</img>';
        copyButton.disabled = true;

        // Opcional: Desabilitar novamente após alguns segundos
        setTimeout(function() {
            copyButton.innerHTML = 'Copiar Link<img src="/static/images/copy.svg" alt="Imagem de Copiar Link">';
            copyButton.disabled = false;
        }, 3000); // 3000 milissegundos (3 segundos)
    } catch (error) {
        console.error('Erro ao copiar link:', error.message);
        alert('Erro ao copiar link: ' + error.message);
    }
}
