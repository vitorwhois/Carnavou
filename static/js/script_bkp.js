
// Aguarda o carregamento completo do DOM antes de executar o script
document.addEventListener('DOMContentLoaded', function () {
    // Função para fazer uma solicitação AJAX usando Fetch
    function obterBlocosPorData() {
        // Substitua a URL pelo caminho correto para o seu manipulador BlocosPorDataHandler
        const url = "/blocosPorData";

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
        })
        .catch(error => {
            // Lida com erros durante a solicitação
            console.error("Erro durante a solicitação:", error.message);
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

    // Função para criar um card com base nos dados do bloco
    function criarCard(bloco) {
        // Aqui, você pode criar a estrutura HTML do card com os dados específicos
        const card = document.createElement('div');
        card.className = 'card';
        card.innerHTML = `
            <h2>${bloco.Nome}</h2>
            <p>${bloco.Local}</p>
            <p>Data: ${bloco.Data}</p>
            <!-- Adicione outras informações do bloco conforme necessário -->
        `;
        return card;
    }

    // Chama a função para obter blocos por data
    obterBlocosPorData();
});
