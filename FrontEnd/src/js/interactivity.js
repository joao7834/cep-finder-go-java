const INPUT = document.getElementById('CEP'); 
const BUTTON = document.getElementById('Botao'); // Só declarando as constantes

INPUT.addEventListener('input', (e) => { 
  let value = e.target.value.replace(/\D/g, ''); // Se não for número não vai
  
  if (value.length > 8) { // Limite de Números
    value = value.slice(0, 8);
  }

  if (value.length > 5) { // Quando passar de 5 add a mascara
    INPUT.value = value.slice(0, 5) + "-" + value.slice(5, 8);
  } else {
    INPUT.value = value; // Se tiver menos de 5 números da pra apagar dboas
  }
});

BUTTON.addEventListener('click', async () => { 
  const CEPCL = INPUT.value.replace('-', '');
  
  if (CEPCL.length !== 8) {
    window.alert("Por favor, insira um CEP valido com 8 digitos");
    return; // Verifica se ta td certinho
  }

 
  const url = `http://localhost:7070/cepR?cep=${CEPCL}`;
  
  try {
    //pesquisa a url
    const AN = await fetch(url);
    
    if (!AN.ok) {
      if (AN.status === 429) {
        window.alert(`SEU IP FOI BLOQUEADO POR RATE LIMIT. CÓDIGO 429`);
      } else {
        window.alert("DEFESA DO SERVIDOR. ALERTA 400 ou CEP INVÁLIDO");
      }
      return;
    }

   // transforma os dados
    const dados = await AN.json();
    
   //mnsg dos dados
    const MEND = `Endereço Encontrado!\n\n` +
                             `Rua: ${dados.logradouro}\n` +
                             `Bairro: ${dados.bairro}\n` +
                             `Cidade: ${dados.localidade} - ${dados.uf}\n` +
                             `DDD: ${dados.ddd}`;

    window.alert(MEND);

  } catch (ERlog) {
    //se der erro vai mostra
    console.error("Erro de conexão com o servidor Go: ", ERlog);
    window.alert("Não foi possível conectar ao servidor de CEP.");
  }
})