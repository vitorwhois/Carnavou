    // Defina a var iável "blocoId" 
    let blocoId;

// Aguarda o carregamento completo do DOM antes de executar o script
document.addEventListener('DOMContentLoaded', function () {


// Função para obter blocos por IDs por url
function obterBlocosPorID() {
    // Obtém os dados codificados da URL da página
    const urlParams = new URLSearchParams(window.location.search);
    const ids = urlParams.get('ids');

    // Verifica se há IDs na URL
    if (!ids) {
    console.error("IDs de blocos não encontrados na URL.");
    return;
}
    // Faz uma solicitação GET para buscar os blocos correspondentes
    fetch('/minhalista?ids=' + ids)
        .then(response => response.json())
        .then(blocos => {
            // Chama a função para adicionar os cards ao container com os blocos recebidos
            adicionarCardsAoContainer(blocos);
        })
        .catch(error => console.error('Erro ao buscar blocos:', error));
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
        <button class="add-list-link" data-bloco-id="${bloco.ID}">Add a lista <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-plus" viewBox="0 0 16 16">
            <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4"></path>
        </svg></button>
    `;
        card.appendChild(cardTitulo);

        // Adiciona a div com as informações do card
        const infoCard = document.createElement('div');
        infoCard.className = 'info-card';

        // Adiciona a div com a data
        const cardData = document.createElement('div');
        cardData.className = 'card-data';
        // Obtém apenas o dia e o mês da propriedade Data
        const diaEMes = bloco.Data.substring(0, 5);

        cardData.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-calendar" viewBox="0 0 16 16">
                <path d="M3.5 0a.5.5 0 0 1 .5.5V1h8V.5a.5.5 0 0 1 1 0V1h1a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h1V.5a.5.5 0 0 1 .5-.5M1 4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4z"></path>
            </svg>
            <span>${diaEMes}</span>
        `;
        infoCard.appendChild(cardData);

        // Adiciona a div com o endereço
        const cardEndereco = document.createElement('div');
        cardEndereco.className = 'card-endereco';
        cardEndereco.innerHTML = `
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-geo-alt" viewBox="0 0 16 16">
        <path d="M12.166 8.94c-.524 1.062-1.234 2.12-1.96 3.07A32 32 0 0 1 8 14.58a32 32 0 0 1-2.206-2.57c-.726-.95-1.436-2.008-1.96-3.07C3.304 7.867 3 6.862 3 6a5 5 0 0 1 10 0c0 .862-.305 1.867-.834 2.94M8 16s6-5.686 6-10A6 6 0 0 0 2 6c0 4.314 6 10 6 10"/>
        <path d="M8 8a2 2 0 1 1 0-4 2 2 0 0 1 0 4m0 1a3 3 0 1 0 0-6 3 3 0 0 0 0 6"/>
      </svg>
            <span>${bloco.Local}</span>
        `;
        infoCard.appendChild(cardEndereco);


        // Adiciona a div com o horário
        const cardHorario = document.createElement('div');
        cardHorario.className = 'card-horario';
        // Somente os 2 primeiros dígitos do horário
        const duasPrimeirasHoras = bloco.Concentracao.substring(0, 2);
        cardHorario.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-clock" viewBox="0 0 16 16">
                <path d="M8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71z"></path>
                <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m7-8A7 7 0 1 1 1 8a7 7 0 0 1 14 0"></path>
            </svg>
            <span>${duasPrimeirasHoras}h</span>
        `;
        infoCard.appendChild(cardHorario);

/*         // Adiciona a div com as categorias - Proxima feature
        const categorias = document.createElement('div');
        categorias.className = 'categorias';
        categorias.textContent = 'Categorias';
        infoCard.appendChild(categorias);

        // Adiciona a div com a lista de categorias (vazia por enquanto)
        const categoriasValor = document.createElement('div');
        categoriasValor.className = 'categorias-valor';
        categoriasValor.innerHTML = '<ul class="d-flex"></ul>';
        infoCard.appendChild(categoriasValor); */

        // Adiciona a div infoCard ao card
        card.appendChild(infoCard);

        return card;
    }

    // Adiciona um ouvinte de evento ao botão de pesquisa
    const botaoPesquisa = document.getElementById('botaoPesquisa');
    if (botaoPesquisa) {
        botaoPesquisa.addEventListener('click', realizarPesquisa);
    }

    // Chama a função para obter blocos por data ao carregar a página
    obterBlocosPorID('');
    chamarMinhaLista();

/*     function adicionarOuvintesDeEventos() {
        var buttons = document.querySelectorAll('.add-list-link');
        buttons.forEach(button => {
            button.addEventListener('click', () => {
                let blocoId = button.getAttribute('data-bloco-id');
                if (blocoId) {
                    console.log(blocoId);
                } else {
                    console.error('O atributo "data-bloco-id" não foi encontrado no botão clicado.');
                }
            });
        });
    } */
    


});




