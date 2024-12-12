// Defina a var iável "blocoId" 
let blocoId;
let blocosSalvos = [];
function salvarBlocoNaLista(blocoId) {
    console.log(blocoId);
    let localSalvos = JSON.parse(localStorage.getItem('blocosSalvos'));
    let blocosSalvos = localSalvos ? localSalvos : []; // Se localSalvos for null, inicialize blocosSalvos como um array vazio
    // Adiciona o novo ID à lista
    if (!blocosSalvos.includes(blocoId)) {
        blocosSalvos.push(blocoId);
        localStorage.setItem('blocosSalvos', JSON.stringify(blocosSalvos)); // Salva blocosSalvos no localStorage, não localSalvos
    }
}

document.getElementById('minhalista').addEventListener('click', ChamaLista);

function ChamaLista() {
    let lista = JSON.parse(localStorage.getItem('blocosSalvos'));
    console.log(blocosSalvos);
    console.log(lista);
    if (lista) {
        window.location.href = `minhalista?ids=${lista.join(',')}`;
    }
    else {
        window.location.href = `minhalista`;
    }
}


// Aguarda o carregamento completo do DOM antes de executar o script
document.addEventListener('DOMContentLoaded', function () {

    const blocosPorPagina = 5;
    let quantidadeBlocosCarregados = 0;

    const btnProximos = document.getElementById('btn-proximos');
    if (btnProximos) {
        btnProximos.addEventListener('click', function () {
            // Carrega mais 5 blocos (ou a quantidade desejada)
            obterBlocosAgenda(obterDataAtualFormatada(), blocosPorPagina, quantidadeBlocosCarregados);
        });
    }
    /*         // Chama a função para obter os primeiros 5 blocos ao carregar a página
            obterBlocosAgenda(blocosPorPagina, quantidadeBlocosCarregados); */

    // Adiciona um ouvinte de evento ao campo de pesquisa
    const campoPesquisa = document.getElementById('nomeInput');
    if (campoPesquisa) {
        campoPesquisa.addEventListener('input', function () {

            // Verifica se o campo de pesquisa está vazio
            if (campoPesquisa.value === '') {
                // Retorna sem enviar a solicitação
                return;
            }
        });
    }

    // Adiciona um ouvinte de evento ao botão de Pesquisa
    const botaoPesquisa = document.getElementById('botaoPesquisa');
    if (botaoPesquisa) {
        botaoPesquisa.addEventListener('click', function () {
            const nomeInput = document.getElementById('nomeInput').value;
            if (nomeInput) {
                pesquisarBlocosPorNome(nomeInput);
            } else {
                console.log('Por favor, insira um nome para pesquisar.');
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
        const url = `/blocos/datas?datas=${dataInput}`;
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
        // BlocosPorNomeHandler
        const url = `/blocos/nomes?nomes=${nomeInput}`;
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
        let buscaErro; // Declarar a variável de erro da busca fora if

        // Verificar se os dados estão vazios
        if (blocosPorData === null || blocosPorData.length === 0) {
            tituloResultado.textContent = 'Opa, não encontrei nenhum bloco 🧐';
            buscaErro = document.createElement('p');
            buscaErro.textContent = 'Não achei nenhum bloco correspondente a sua busca. Você pode tentar outra busca mudando os valores ou conferir a agenda do dia aqui:';
            //Alterar o estilo do tituloResultado
            tituloResultado.style.marginBottom = '8px';
        } else {
            tituloResultado.textContent = 'Resultados da busca';
        }

        blocosContainer.appendChild(tituloResultado);

        // Adicionar buscaErro ao contêiner apenas se estiver definido
        if (buscaErro) {
            blocosContainer.appendChild(buscaErro);
        }

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
            <button class="add-list-link btn-card" data-bloco-id="${bloco.ID}">Salvar<img src="/static/images/plus.svg" alt="Soma Icon" width="24" height="24"</img>
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
        <img src="/static/images/calendar.svg" alt="Icon Calendario" width="24" height="24"
    </img>
            <span>${diaEMes}</span>
        `;
        infoCard.appendChild(cardData);

        // Adiciona a div com o horário
        const cardHorario = document.createElement('div');
        cardHorario.className = 'card-horario col-6';
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

    // Adiciona um ouvinte de evento ao botão de Filtros
    const botaoFiltrar = document.getElementById('botaoFiltrar');
    if (botaoFiltrar) {
        botaoFiltrar.addEventListener('click', filtrar);
    }

    // Função para realizar a pesquisa com base no local e data
    function filtrar() {
        var dataInput = document.getElementById('dataInput').value;
        var localInput = document.getElementById('localInput').value;

        // Verificar se há dados nos campos de data e local
        if (dataInput && localInput) {
            // Se ambos os campos estiverem preenchidos, realizar pesquisa por data e local
            ObterBlocosPorDataESubprefeitura(dataInput, localInput);
        } else if (dataInput) {
            // Se apenas o campo de data estiver preenchido, realizar pesquisa por data
            obterBlocosPorData(dataInput);
        } else if (localInput) {
            // Se apenas o campo de local estiver preenchido, realizar pesquisa por local
            obterBlocosPorLocal(localInput);
        } else {
            // Se nenhum campo estiver preenchido, exibir uma mensagem ou tomar outra ação
            // Falta chamar tratamento de erro
            console.log('Nenhum critério de pesquisa fornecido.');
        }
    }
    function ObterBlocosPorDataESubprefeitura(dataInput, localInput) {
        const url = `/blocos/filtros?datas=${encodeURIComponent(dataInput)}&subprefeituras=${encodeURIComponent(localInput)}`;
        console.log(dataInput, localInput);
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
                return response.json();
            })
            .then(data => {
                adicionarCardsAoContainer(data);
                adicionarOuvintesDeEventos();
            })
            .catch(error => {
                console.error('Erro:', error);
            });
    }

    function obterBlocosPorLocal(localInput) {
        const url = `/subprefeitura?subprefeitura=${localInput}`;
        console.log(localInput);
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
                return response.json();
            })
            .then(data => {
                adicionarCardsAoContainer(data);
                adicionarOuvintesDeEventos();
            })
            .catch(error => {
                console.error('Erro:', error);
            });
    }




    function adicionarOuvintesDeEventos() {
        var buttons = document.querySelectorAll('.add-list-link');
        buttons.forEach(button => {
            button.addEventListener('click', () => {
                let blocoId = button.getAttribute('data-bloco-id');
                salvarBlocoNaLista(blocoId);

                // Altera o texto do botão e desabilita
                button.innerHTML = 'Salvo <img src="/static/images/check.svg" alt="Check Icon" width="24" height="24"</img>';
                button.disabled = true;

                salvarBlocoNaLista(blocoId);
            });
        });

    }



    // Adiciona a div para os cards abaixo da div agenda
    const agendaContainer = document.querySelector('.agenda');
    const cardsContainer = document.createElement('div');
    cardsContainer.className = 'cards-container';
    agendaContainer.appendChild(cardsContainer);

    // Função para obter a data atual no formato DD/MM/YYYY
    function obterDataAtualFormatada() {
        let dataAtual = new Date();
        let dia = String(dataAtual.getDate()).padStart(2, '0');
        let mes = String(dataAtual.getMonth() + 1).padStart(2, '0');
        let ano = dataAtual.getFullYear();

        // Cria uma data no formato DD/MM/YYYY
        let dataFormatada = `${dia}/${mes}/${ano}`;

        // Cria uma data para 18/02/2024
        let dataLimite = new Date(2024, 1, 18);

        // Compara a data atual com a data limite
        if (dataAtual > dataLimite) {
            return '18/02/2024';
        } else {
            return dataFormatada;
        }
    }

    // Define a quantidade e o startIndex
    let quantidade = 5;
    let startIndex = 0;

    // Chama a função para obter blocos por data ao carregar a página
    obterBlocosAgenda(obterDataAtualFormatada(), quantidade, startIndex);

    function adicionarCardsNaAgenda(cards) {
        // Adiciona os 5 cards ao contêiner
        cards.forEach(bloco => {
            const card = criarCard(bloco);
            cardsContainer.appendChild(card);
        });
    }

    function obterBlocosAgenda(data, quantidade, startIndex) {
        const url = `/blocos/datas?datas=${data}`;
        console.log(url);
        console.log(data);

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
                if (data === null) {
                    console.error('Os dados recebidos são nulos.');
                    return;
                }

                // Verifica se há blocos suficientes para carregar
                if (startIndex < data.length) {
                    // Chama a função para adicionar os 5 cards ao contêiner
                    adicionarCardsNaAgenda(data.slice(startIndex, startIndex + quantidade));
                    quantidadeBlocosCarregados += quantidade;
                    adicionarOuvintesDeEventos();
                } else {
                    console.log('Não há mais blocos para carregar.');
                }
            })
            .catch(error => {
                // Lida com erros durante a solicitação
                console.error("Erro durante a solicitação:", error.message);
            });
    }

    document.getElementById("botaoLimpar").addEventListener("click", function () {
        document.getElementById("nomeInput").value = "";
        document.getElementById("dataInput").value = "";
        document.getElementById("localInput").value = "";
    });



});






