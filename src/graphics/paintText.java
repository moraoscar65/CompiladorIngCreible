/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package graphics;

import java.util.ArrayList;
import java.util.Arrays;



public class paintText{
    
    private static ArrayList<String> palabrasReservadas = new ArrayList();

    public paintText() {    
        palabrasReservadas.add("program");
        palabrasReservadas.add("if");
        palabrasReservadas.add("else");
        palabrasReservadas.add("fi");
        palabrasReservadas.add("do");
        palabrasReservadas.add("until");
        palabrasReservadas.add("while");
        palabrasReservadas.add("read");
        palabrasReservadas.add("write");
        palabrasReservadas.add("float");
        palabrasReservadas.add("int");
        palabrasReservadas.add("bool");
        palabrasReservadas.add("not");
        palabrasReservadas.add("and");
        palabrasReservadas.add("or");
        palabrasReservadas.add("{");
        palabrasReservadas.add("}");
        palabrasReservadas.add(";");
        palabrasReservadas.add("then");
    }

    
    public ArrayList<String> separarTexto(String textoCompleto){
        ArrayList<String> texto = new ArrayList<String>();
        String[] textoLineado = separarPorN(textoCompleto);
        String[] textoEspaciado;
        
        for(String linea:textoLineado){
            textoEspaciado = separarPorS( linea );
            texto.addAll(Arrays.asList( textoEspaciado ));
        }
       
        /*texto.stream().forEach((b) -> {
            System.out.print("["+b+"]");
        });*/
        
        //System.out.println("");
        
        return texto;
    }
    
    public String[] separarPorS(String cadena){
        ArrayList<String> text = new ArrayList<String>();
        String[] separado = cadena.split("\\s");
        
        for(String a:separado){
            text.add(a);
            text.add(" ");
        }
        try{
            text.remove(text.size()-1);
        }
        catch(Exception e){
            System.out.println(e);
            text.add("\n");
        }
        return text.toArray(new String[0]);
    }
    
    public String[] separarPorN(String cadena){
        ArrayList<String> text = new ArrayList<String>();
        String[] separado = cadena.split("\\n");
        
        for(String a:separado){
            text.add(a);
            text.add("\n");
        }
        text.remove(text.size()-1);
        
        return text.toArray(new String[0]);
    }
        
    public static boolean checkString(String word){
        return palabrasReservadas.stream().anyMatch((PR) -> (word.equals(PR)));
    }
    
    public static boolean checkKeyValid(int value){
        return value != 32 && value != 10  && value != 40 && value != 39 && value != 38 && value != 37 && value != 9 && value != 20;
    }
}
