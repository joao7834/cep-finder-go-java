const INPUT = document.getElementById('CEP'); //Constante da caixa de txt
const BUTTON = document.getElementById('Botao'); //Constante do bofao

INPUT.addEventListener('input', (e) => { //Quando o usuario coloca/tira um caractere, acontece essa funçâo
  let value = e.target.value.replace(/\D/g, ''); //Tira oqq nn for numero
  
  if (value.length > 8) { //Limita o CEP em 8 numeros
    value = value.slice(0,8);
  }

  if (value.length > 5) { //Se o usuario digitar + de 5 numeros, add a mascara mo cep
      INPUT.value = value.slice(0, 5) + "-" + value.slice(5,8);
  } else {
      INPUT.value = value //Se tiver menos de 5 numeros, deixa apagar dboas
  }
})

BUTTON.addEventListener('click', () => { //Quando o usuário apertar, dará um aviso
  window.alert(INPUT.value);
});