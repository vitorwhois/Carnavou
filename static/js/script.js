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
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-calendar" viewBox="0 0 16 16">
                <path d="M3.5 0a.5.5 0 0 1 .5.5V1h8V.5a.5.5 0 0 1 1 0V1h1a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h1V.5a.5.5 0 0 1 .5-.5M1 4v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4z"></path>
            </svg>
            <span>${diaEMes}</span>
        `;
        infoCard.appendChild(cardData);
    
        // Adiciona a div com o horário
        const cardHorario = document.createElement('div');
        cardHorario.className = 'card-horario col-6';
        const duasPrimeirasHoras = bloco.Concentracao.substring(0, 2);
        cardHorario.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-clock" viewBox="0 0 16 16">
                <path d="M8 3.5a.5.5 0 0 0-1 0V9a.5.5 0 0 0 .252.434l3.5 2a.5.5 0 0 0 .496-.868L8 8.71z"></path>
                <path d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16m7-8A7 7 0 1 1 1 8a7 7 0 0 1 14 0"></path>
            </svg>
            <span>${duasPrimeirasHoras}h</span>
        `;
        infoCard.appendChild(cardHorario);
    
        // Adiciona a div com o endereço
        const cardEndereco = document.createElement('div');
        cardEndereco.className = 'card-endereco';
        cardEndereco.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-geo-alt" viewBox="0 0 16 16">
                <path d="M12.166 8.94c-.524 1.062-1.234 2.12-1.96 3.07A32 32 0 0 1 8 14.58a32 32 0 0 1-2.206-2.57c-.726-.95-1.436-2.008-1.96-3.07C3.304 7.867 3 6.862 3 6a5 5 0 0 1 10 0c0 .862-.305 1.867-.834 2.94M8 16s6-5.686 6-10A6 6 0 0 0 2 6c0 4.314 6 10 6 10"></path>
                <path d="M8 8a2 2 0 1 1 0-4 2 2 0 0 1 0 4m0 1a3 3 0 1 0 0-6 3 3 0 0 0 0 6"></path>
            </svg>
            <span>${bloco.Local}</span><br>
            <p>${bloco.Subprefeitura}</p>
        `;
       
    
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
      


/* function chamarMinhaLista() {
    // Codifica a lista de blocos
    let dadosCodificados = btoa(JSON.stringify(blocosSalvos));
  
    // Cria o link para a lista do usuário
    let link = `/minhalista/${dadosCodificados}`;
  
    // Chama a rota /minhalista com os dados codificados
    fetch(link)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        // Aqui você pode adicionar a lógica para criar os cards com os dados recebidos
        adicionarCardsAoContainer(data);
      })
      .catch(error => console.error('Erro:', error));
  } */

});

//Sidebar





