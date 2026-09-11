public class Main {
    public static void main(String[] args) {
        try {
            Refator app = new Refator();
            app.IniciarServerJava(); // <--- Sobbe o servidor do Davi
        } catch (Exception e) {
            System.err.println("Erro ao iniciar o Java: " + e.getMessage());
        }
    }
}