    // Defina a var iável "blocoId" 
    let blocoId;
    let blocosSalvos = [];

function salvarBlocoNaLista(blocoId) {
  console.log(blocoId);

  // Adiciona o novo ID à lista
  if (!blocosSalvos.includes(blocoId)) {
    blocosSalvos.push(blocoId);
  }

  // Imprime o valor de blocosSalvos no console
  console.log(blocosSalvos);
}


function ChamaLista (){
    let lista = blocosSalvos;
    console.log(blocosSalvos);
    console.log(lista);
    window.location.href = `minhalista?ids=${blocosSalvos.join(',')}`;
}


// Aguarda o carregamento completo do DOM antes de executar o script
document.addEventListener('DOMContentLoaded', function () {

  // Adiciona um ouvinte de evento ao campo de pesquisa
  const campoPesquisa = document.getElementById('nomeInput');
  if (campoPesquisa) {
    campoPesquisa.addEventListener('input', function () {
      // Imprime o termo de pesquisa
      console.log('Termo de pesquisa:', campoPesquisa.value);

      // Verifica se o campo de pesquisa está vazio
      if (campoPesquisa.value === '') {
        // Retorna sem enviar a solicitação
        return;
      }
    });
  }

    // Função para realizar a pesquisa com base nos campos preenchidos
    function realizarPesquisa() {
        var nomeInput = document.getElementById('nomeInput').value;
        var dataInput = document.getElementById('dataInput').value;

        // Verificar se há dados nos campos de nome e data
        if (nomeInput || dataInput) {
            // Se pelo menos um dos campos estiver preenchido, realizar a pesquisa
            if (nomeInput && dataInput) {
                // Se ambos estiverem preenchidos, realizar pesquisa por ambos
                pesquisarBlocosPorNome(nomeInput);
                obterBlocosPorData(dataInput);
            } else if (nomeInput) {
                // Se apenas o campo de nome estiver preenchido, realizar pesquisa por nome
                pesquisarBlocosPorNome(nomeInput);
            } else if (dataInput) {
                // Se apenas o campo de data estiver preenchido, realizar pesquisa por data
                obterBlocosPorData(dataInput);
            }
        } else {
            // Se nenhum campo estiver preenchido, exibir uma mensagem ou tomar outra ação
            console.log('Nenhum critério de pesquisa fornecido.');
        }
    }

    // Função para fazer uma solicitação AJAX usando Fetch
    function obterBlocosPorData(dataInput) {
        // Substitua a URL pelo caminho correto para o seu manipulador BlocosPorDataHandler
        const url = `/blocosPorData?data=${dataInput}`;
        console.log(dataInput);
        console.log(url);

        // Faz uma solicitação GET para o servidor
        fetch(url, {
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
        .then(data => {
            // Manipula os dados recebidos (no formato JSON)
            console.log("Blocos recebidos:", data);

            // Chama a função para adicionar os cards ao container com os dados recebidos
            adicionarCardsAoContainer(data);
            adicionarOuvintesDeEventos();
        })
        .catch(error => {
            // Lida com erros durante a solicitação
            console.error("Erro durante a solicitação:", error.message);
        });
    }

    // Função para pesquisar blocos por nome
    function pesquisarBlocosPorNome(nomeInput) {
        // Substitua a URL pelo caminho correto para o seu manipulador BlocosPorNomeHandler
        const url = `/pesquisarBlocos?nome=${nomeInput}`;
        console.log(nomeInput);
        console.log(url);
        // Faz uma solicitação GET para o servidor
        fetch(url, {
            method: 'GET',
            headers: {
                'X-Requested-With': 'XMLHttpRequest',
            },
        })
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erro na solicitação: ${response.status}`);
            }
    
            // Verifica se a resposta é um JSON válido
            const contentType = response.headers.get('content-type');
            if (contentType && contentType.includes('application/json')) {
                // Se for JSON, converte a resposta para JSON
                return response.json();
            } else {
                // Se não for JSON, lança um erro indicando uma resposta inesperada
                throw new Error('Resposta inesperada do servidor: não é JSON');
            }
        })
        .then(data => {
            console.log("Blocos pesquisados:", data);

            // Chama a função para adicionar os cards ao container com os dados da pesquisa
            adicionarCardsAoContainer(data);
            adicionarOuvintesDeEventos();
        })
        .catch(error => {
            console.error("Erro durante a pesquisa:", error.message);
        });
    }

    // Função para adicionar os cards ao contêiner com base nos dados recebidos
    function adicionarCardsAoContainer(blocosPorData) {
        const blocosContainer = document.getElementById('blocosContainer');

     // Limpar qualquer conteúdo existente no contêiner
     blocosContainer.innerHTML = '';

     // Adicionar o h2 com o texto "Resultados da busca"
     const tituloResultado = document.createElement('h2');
     tituloResultado.textContent = 'Resultados da busca';
     blocosContainer.appendChild(tituloResultado);

        // Adicionar os cards ao contêiner
        blocosPorData.forEach(bloco => {
            const card = criarCard(bloco);
            blocosContainer.appendChild(card);
        });
    }

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
            <button class="add-list-link btn-card" data-bloco-id="${bloco.ID}">Salvar
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-plus" viewBox="0 0 16 16">
                    <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4"></path>
                </svg>
            </button>
        `;
        card.appendChild(cardTitulo);
    
        // Adiciona a div com as informações do card
        const infoCard = document.createElement('div');
        infoCard.className = 'info-card';
    
        // Adiciona a div com a data
        const cardData = document.createElement('div');
        cardData.className = 'card-data col-6';
        const diaEMes = bloco.Data.substring(0, 5);
        cardData.innerHTML = `
        <img src="/static/images/Calendar.svg" alt="Icon Calendario" width="24" height="24"
    </img>
            <span>${diaEMes}</span>
        `;
        infoCard.appendChild(cardData);
    
        // Adiciona a div com o horário
        const cardHorario = document.createElement('div');
        cardHorario.className = 'card-horario col-6';
        const duasPrimeirasHoras = bloco.Concentracao.substring(0, 2);
        cardHorario.innerHTML = `
        <img src="/static/images/Clock.svg" alt="Icon Relogio" width="24" height="24"
    </img>
            <span>${duasPrimeirasHoras}h</span>
        `;
        infoCard.appendChild(cardHorario);
    
        // Adiciona a div com o endereço
        const cardEndereco = document.createElement('div');
        cardEndereco.className = 'card-endereco endereco-column';
        cardEndereco.innerHTML = `
        <img src="/static/images/Map_Pin.svg" alt="Map Pin Icon" width="24" height="24"
    </img>
    <div class="endereco-content">
    <a href="https://www.google.com/maps/search/${bloco.Local}" target="_blank">${bloco.Local}</a>
    <p>${bloco.Subprefeitura}</p>
</div>
`;
       
    //<span>${bloco.Local}A</span>
        // Adiciona a div infoCard ao card
        card.appendChild(infoCard);
        card.appendChild(cardEndereco);
        return card;
    }
    
    

    // Adiciona um ouvinte de evento ao botão de pesquisa
    const botaoPesquisa = document.getElementById('botaoPesquisa');
    if (botaoPesquisa) {
        botaoPesquisa.addEventListener('click', realizarPesquisa);
    }
     // Adiciona um ouvinte de evento ao botão de Filtros
    const botaoFiltrar = document.getElementById('botaoFiltrar');
    if (botaoFiltrar) {
    botaoFiltrar.addEventListener('click', realizarPesquisa);
    }

    // Chama a função para obter blocos por data ao carregar a página
    obterBlocosPorData('');

    function adicionarOuvintesDeEventos() {
        var buttons = document.querySelectorAll('.add-list-link');
        buttons.forEach(button => {
            button.addEventListener('click', () => {
                let blocoId = button.getAttribute('data-bloco-id');
                salvarBlocoNaLista(blocoId);
            });
        });
    }
      


    // Adiciona a div para os cards abaixo da div agenda
    const agendaContainer = document.querySelector('.agenda');
    const cardsContainer = document.createElement('div');
    cardsContainer.className = 'cards-container';
    agendaContainer.appendChild(cardsContainer);

    // Chama a função para obter blocos por data ao carregar a página
    obterBlocosAgenda('03/02/2024');

    function adicionarCardsNaAgenda(cards) {
        // Adiciona os 5 cards ao contêiner
        cards.forEach(bloco => {
            const card = criarCard(bloco);
            cardsContainer.appendChild(card);
        });
    }

    function obterBlocosAgenda() {
        // Substitua a URL pelo caminho correto para o seu manipulador BlocosPorDataHandler
        const url = `/blocosPorData?data=03/02/2024`;
        console.log();
        console.log(url);

        // Faz uma solicitação GET para o servidor
        fetch(url, {
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
        .then(data => {
            // Manipula os dados recebidos (no formato JSON)
            console.log("Blocos recebidos:", data);

            // Chama a função para adicionar os cards ao container com os dados recebidos
            adicionarCardsNaAgenda(data);
            adicionarOuvintesDeEventos();
        })
        .catch(error => {
            // Lida com erros durante a solicitação
            console.error("Erro durante a solicitação:", error.message);
        });
    }
});

//Sidebar





