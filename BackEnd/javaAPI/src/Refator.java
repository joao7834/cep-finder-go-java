import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.URI;
import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

public class Refator{

private final HttpClient client = HttpClient.newHttpClient();

    public void IniciarServerJava () throws IOException {

        HttpServer ServerJava = HttpServer.create(new InetSocketAddress(9090), 0);
        ServerJava.createContext("/cepR/", new HttpHandler(){
            @Override

        public void handle (HttpExchange exchange) throws IOException{

            try {
                URI completUrl = exchange.getRequestURI();
                String query = completUrl.getQuery();
                String cep = "";

                if (query != null) {
                    String[] pares = query.split("&");
                for(String par : pares){
                    
                    int indx = par.indexOf("=");

                    if (indx != -1) {
                    String chave = URLDecoder.decode(par.substring(0,indx),StandardCharsets.UTF_8);

                    if (chave.equalsIgnoreCase("cep")) {
                        cep = URLDecoder.decode(par.substring(indx + 1), StandardCharsets.UTF_8);
                        break;
                    }
                }
            }    
        }    
            cep =cep.replaceAll("[^0-9]","").trim();
            System.out.print("Carregando...");

            String jsonViaCep = buscarNoViaCep(cep);

            System.out.println("Resposta do Json: " + jsonViaCep);

             if (jsonViaCep.contains("\"erro\": \"true\"") || jsonViaCep.contains("\"erro\":true")) {
                        jsonViaCep = "{\"status\": \"bloqueio_viacep\", \"mensagem\": \"O ViaCEP nao encontrou esse CEP ou barrou o IP.\"}";
                        exchange.getResponseHeaders().set("Content-Type", "application/json");
                        exchange.sendResponseHeaders(429, jsonViaCep.getBytes(StandardCharsets.UTF_8).length);
                    } else {
                        exchange.getResponseHeaders().set("Content-Type", "application/json");
                        exchange.sendResponseHeaders(200, jsonViaCep.getBytes(StandardCharsets.UTF_8).length);
                    }
                    
                    OutputStream os = exchange.getResponseBody();
                    os.write(jsonViaCep.getBytes(StandardCharsets.UTF_8));
                    os.close();

                    System.out.println("Dados retornados");

                } catch (Exception e) {
                    String erro = "{\"erro\": \"Erro interno no Java: " + e.getMessage() + "\"}";
                    exchange.getResponseHeaders().set("Content-Type", "application/json");
                    exchange.sendResponseHeaders(500, erro.getBytes(StandardCharsets.UTF_8).length);
                    exchange.getResponseBody().write(erro.getBytes(StandardCharsets.UTF_8));
                    exchange.getResponseBody().close();
                }
            }
        });

        ServerJava.setExecutor(null);
        ServerJava.start();
        System.out.println("O Servidor do Java tá no 9090");
    }

    private String buscarNoViaCep(String cep) throws IOException, InterruptedException {
        String urlFinal = "https://viacep.com.br/ws/" + cep + "/json/";
        
        HttpRequest requestViaCep = HttpRequest.newBuilder()
                .uri(URI.create(urlFinal))
                .header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)") 
                .GET()
                .build();

        HttpResponse<String> response = client.send(requestViaCep, HttpResponse.BodyHandlers.ofString());
        return response.body();

    }
}
    
        

        


    
