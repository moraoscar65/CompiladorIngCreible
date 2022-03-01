/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package functions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.swing.JFileChooser;


public class filesAdmin {
    private static JFileChooser fileChooser;
    
    public filesAdmin() {
        fileChooser = new JFileChooser();
    }
    
    
    // abre el explorador de archivos y nos devuelve la ruta al archivo que seleccionemos.
    public String seleccionarArchivo(){
        fileChooser.setCurrentDirectory(new File(System.getProperty("user.home")));
        String x = "No se selecciono ningun archivo.";
        int result = fileChooser.showOpenDialog(null);
        if (result == JFileChooser.APPROVE_OPTION) {
            File selectedFile = fileChooser.getSelectedFile();
            x = selectedFile.getAbsolutePath();
        }
                
        return x;
    }
    
    // abre el explorador de archivos y nos devuelve la ruta a la carpeta que seleccionemos.
    public String seleccionarCarpeta(){
        fileChooser.setCurrentDirectory(new File(System.getProperty("user.home")));
        String x = "No se selecciono ningun archivo.";
        int result = fileChooser.showSaveDialog(null);
        if (result == JFileChooser.APPROVE_OPTION) {
            File selectedFile = fileChooser.getCurrentDirectory();
            x = selectedFile.getAbsolutePath();
        }
                
        return x;
    }
  
    public void guardarTxt(String texto, String ubicacion){
        FileWriter fichero = null;
        PrintWriter pw = null;
        try {
            fichero = new FileWriter(ubicacion);
            pw = new PrintWriter(fichero);
            pw.println(texto);
        } catch (Exception i) { i.printStackTrace();
        } finally {
            try {  // Nuevamente aprovechamos el finally para asegurarnos que se cierra el fichero.
                if (null != fichero) fichero.close();
            } catch (Exception e2) {e2.printStackTrace(); }
        }
    }
    
    public String leerTxt(String ubicacion){
        String texto = "";
        String cadena = "";
        FileReader f;
        try {
            f = new FileReader(ubicacion);
            BufferedReader b = new BufferedReader(f);
            while((cadena = b.readLine())!=null) {
                texto += cadena + "\n";
            }
            b.close();
        } catch (FileNotFoundException ex) {
            Logger.getLogger(filesAdmin.class.getName()).log(Level.SEVERE, null, ex);
            System.out.println("ERROR AL LEER ARCHIVO - "+ubicacion);
        } catch (IOException ex) {
            Logger.getLogger(filesAdmin.class.getName()).log(Level.SEVERE, null, ex);
            System.out.println("ERROR AL LEER ARCHIVO - "+ubicacion);
        }
        
        return texto;
    }
    
    
}
