package graphics;

import functions.ejecutadorGO;
import functions.filesAdmin;
import functions.funciones;
import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Dimension;
import java.awt.Font;
import java.awt.Toolkit;
import java.awt.event.ActionEvent;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;
import java.io.File;
import java.util.Scanner;
import java.util.ArrayList;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.swing.GroupLayout;
import javax.swing.Icon;
import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JMenu;
import javax.swing.JMenuBar;
import javax.swing.JMenuItem;
import javax.swing.JOptionPane;
import javax.swing.JPanel;
import javax.swing.JScrollPane;
import javax.swing.JTextPane;
import javax.swing.JToolBar;
import javax.swing.SwingUtilities;
import javax.swing.UIManager;
import javax.swing.event.DocumentEvent;
import javax.swing.event.DocumentListener;
import javax.swing.text.BadLocationException;
import javax.swing.text.Document;
import javax.swing.text.SimpleAttributeSet;
import javax.swing.text.StyleConstants;


public class CompiladorFrame {

    private static int hei;
    private static int wid;
    private static JFrame pantalla;
    private static JTextPane errorsArea;
    private static JTextPane codeArea;
    private static JTextPane analizadorArea;
    // variables de diseno
    private static String letra_Fuente;
    private static Color texto_color;
    private static Color identificador_color;
    private static int texto_size;
    private static int analizador_size;
    private static int errors_size;
    private static Color back_Color;
    private static boolean EditedFile;
    private static int cursorPosition;
    private static filesAdmin fa = new filesAdmin();
    // variables de texto de paneles
    private static String lexicoTxt;
    private static String sintacticoTxt;
    private static String semanticoTxt;
    private static String codigoInterTxt;
    
    private static ArrayList<String> palabrasReservadas = new ArrayList();//

    //contructor
    public CompiladorFrame() {
        palabrasReservadas.add("program");palabrasReservadas.add("if");palabrasReservadas.add("else");
        palabrasReservadas.add("fi");palabrasReservadas.add("do");palabrasReservadas.add("until");palabrasReservadas.add("while");
        palabrasReservadas.add("read");palabrasReservadas.add("write");palabrasReservadas.add("float");palabrasReservadas.add("int");
        palabrasReservadas.add("bool");palabrasReservadas.add("not");palabrasReservadas.add("and");palabrasReservadas.add("or");
        
        Dimension screenSize = Toolkit.getDefaultToolkit().getScreenSize();
        hei = screenSize.height; // 1080
        wid = screenSize.width;  // 1920
        pantalla = new JFrame("Compilador");
        codeArea = new JTextPane();
        errorsArea = new JTextPane();
        analizadorArea = new JTextPane();
        letra_Fuente = "consolas";
        identificador_color = Color.RED;
        texto_color = Color.LIGHT_GRAY;
        back_Color = Color.darkGray;
        texto_size = 16;
        analizador_size = 14;
        errors_size = 16;
        EditedFile = false;
        agregarContenidoFrame();
        limpiarArchivos();
    }

    //contructor con argumentos
    public CompiladorFrame(String font, Color text, Color identifier, Color background, int textSize, int analizerSize, int errorsSize) {
        Dimension screenSize = Toolkit.getDefaultToolkit().getScreenSize();
        hei = screenSize.height; // 1080
        wid = screenSize.width;  // 1920
        pantalla = new JFrame("Compiladores");
        codeArea = new JTextPane();
        errorsArea = new JTextPane();
        analizadorArea = new JTextPane();
        letra_Fuente = font;
        identificador_color = text;
        texto_color = identifier;
        back_Color = background;
        texto_size = textSize;
        analizador_size = analizerSize;
        errors_size = errorsSize;
        EditedFile = false;
        agregarContenidoFrame();
        limpiarArchivos();
    }

    public void mostrarFrame() {
        //pantalla.setIconImage(Toolkit.getDefaultToolkit().getImage(getClass().getResource("../recursos/potato_icon.png")));
        pantalla.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        pantalla.setMinimumSize(new Dimension(wid - 850, hei - 150));
        pantalla.setExtendedState(JFrame.MAXIMIZED_BOTH);
        pantalla.setResizable(true);
        pantalla.setVisible(true);
    }

    // agrega todos los componentes al frame principal
    private void agregarContenidoFrame() {
        pantalla.add(crearPanelP());
        pantalla.getContentPane().add(agregarBarraMenu(), BorderLayout.NORTH);
    }

    //crear el panel principal
    private static JPanel crearPanelP() {
        funciones f = new funciones();
        int paddingWID = f.porcentaje(wid, 1);
        int paddingHEI = f.porcentaje(hei, 3);

        JPanel panel = new JPanel();
        panel.setBackground(back_Color);
        GroupLayout panelLayout = new javax.swing.GroupLayout(panel);
        panel.setLayout(panelLayout);
        panelLayout.setHorizontalGroup(
                panelLayout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                .addGap(0, 795, Short.MAX_VALUE)
        );
        panelLayout.setVerticalGroup(
                panelLayout.createParallelGroup(javax.swing.GroupLayout.Alignment.LEADING)
                .addGap(0, 199, Short.MAX_VALUE)
        );

        //....  crear los TextArea, Scrollbars y Tool bars del el IDE
        codeArea = new JTextPane();
        codeArea.setForeground(texto_color);
        codeArea.setCaretColor(Color.WHITE);
        
        codeArea.setSelectionColor(Color.white);
        codeArea.setBackground(Color.DARK_GRAY);
        codeArea.setFont(new Font(Font.SANS_SERIF, Font.PLAIN, texto_size));
        cursorPosition = codeArea.getText().length();
       
        JScrollPane scrollCodeArea = new JScrollPane(codeArea);
        scrollCodeArea.setVerticalScrollBarPolicy(JScrollPane.VERTICAL_SCROLLBAR_ALWAYS);
        scrollCodeArea.setHorizontalScrollBarPolicy(JScrollPane.HORIZONTAL_SCROLLBAR_ALWAYS);
        scrollCodeArea.setLocation((2 * paddingWID), paddingHEI);
        scrollCodeArea.setSize(wid - f.porcentaje(wid, 60), hei - f.porcentaje(hei, 35));
        
        codeArea.addKeyListener(new KeyListener() {
            @Override
            public void keyTyped(KeyEvent e) {}
            @Override
            public void keyPressed(KeyEvent e) { }
            @Override
            public void keyReleased(KeyEvent e) {
                codeArea.setEditable(false);
                paintText paint = new paintText();
                Document doc = codeArea.getDocument();
                ArrayList<String> textoSeparado;
                try {
                    if(paintText.checkKeyValid(e.getKeyCode())){
                        textoSeparado = paint.separarTexto( doc.getText(0, doc.getLength()) );
                        codeArea.setText("");
                        textoSeparado.stream().forEach((palabra) -> { insertarPalabra(palabra); });
                    }
                } catch (BadLocationException ex) { Logger.getLogger(CompiladorFrame.class.getName()).log(Level.SEVERE, null, ex); }
                codeArea.setEditable(true);
            }
            
            public void insertarPalabra(String palabra){
                Runnable insert = new Runnable() {
                    @Override
                    public void run() {
                        SimpleAttributeSet attrs = new SimpleAttributeSet();
                        if(paintText.checkString(palabra))StyleConstants.setForeground(attrs, Color.RED);
                        else StyleConstants.setForeground(attrs, Color.LIGHT_GRAY);
                        try {
                            codeArea.getStyledDocument().insertString(codeArea.getStyledDocument().getLength(), palabra, attrs);
                        } catch (BadLocationException ex) {
                            Logger.getLogger(CompiladorFrame.class.getName()).log(Level.SEVERE, null, ex);
                        }
                    }
                };       
                SwingUtilities.invokeLater(insert);
            }

        });

        JToolBar toolBarAnalizador = agregarBarraHerramientas();
        toolBarAnalizador.setLocation(wid - f.porcentaje(wid, 61) + (4 * paddingWID), paddingHEI);
        toolBarAnalizador.setSize(wid - f.porcentaje(wid, 43), paddingHEI);

        analizadorArea = new JTextPane();
        analizadorArea.setText("Presisona 'Compilar' para comezar");
        analizadorArea.setForeground(Color.white);
        analizadorArea.setBackground(Color.DARK_GRAY);
        analizadorArea.setFont(new Font(Font.SANS_SERIF, Font.PLAIN, analizador_size));
        
        
        analizadorArea.setEditable(false);
        JScrollPane scrollAnalizadorArea = new JScrollPane(analizadorArea);
        
        scrollAnalizadorArea.setVerticalScrollBarPolicy(JScrollPane.VERTICAL_SCROLLBAR_ALWAYS);
        scrollAnalizadorArea.setHorizontalScrollBarPolicy(JScrollPane.HORIZONTAL_SCROLLBAR_ALWAYS);
        scrollAnalizadorArea.setLocation(wid - f.porcentaje(wid, 61) + (4 * paddingWID), (2 * paddingHEI));
        scrollAnalizadorArea.setSize(wid - f.porcentaje(wid, 43), hei - f.porcentaje(hei, 35) - paddingHEI);
        scrollAnalizadorArea.setAutoscrolls(true);
        scrollAnalizadorArea.setWheelScrollingEnabled(true);

        errorsArea = new JTextPane();
        errorsArea.setText("\n Compis");
        errorsArea.setBackground(Color.DARK_GRAY);
        errorsArea.setForeground(texto_color);
        errorsArea.setFont(new Font("Arial", Font.PLAIN, errors_size));
        errorsArea.setEditable(false);
        JScrollPane scrollErrorsArea = new JScrollPane(errorsArea);
        scrollErrorsArea.setVerticalScrollBarPolicy(JScrollPane.VERTICAL_SCROLLBAR_ALWAYS);
        scrollErrorsArea.setLocation((2 * paddingWID), hei - f.porcentaje(hei, 28));
        scrollErrorsArea.setSize(wid - (3 * paddingWID), f.porcentaje(hei, 15));

        //.. agrega los componenetes al Panel ...
        panel.add(scrollCodeArea, BorderLayout.PAGE_START);
        panel.add(scrollErrorsArea, BorderLayout.PAGE_START);
        panel.add(toolBarAnalizador, BorderLayout.PAGE_START);
        panel.add(scrollAnalizadorArea, BorderLayout.PAGE_START);

        return panel;
    }

    // agrega el contenido del Menu principal y lo retorna listo para agregarlo al panel
    private static JMenuBar agregarBarraMenu() {
        ArrayList<String> palabras = new ArrayList<>();
        JMenuBar menu = new JMenuBar();

        palabras.add("Archivo");
        palabras.add("Editar");
        palabras.add("Opciones");
        palabras.add("Ayuda");
        palabras.add("Compilar");
        palabras.stream().forEach((temp) -> {
            menu.add(agregarMenuItem(temp));
        });

        return menu;
    }

    // Agrega los JMenu al JMenuBar
    private static JMenu agregarMenuItem(String a) {
        JMenu b = new JMenu(a);
        b.setBorderPainted(false);
        b.setFont(new Font("Arial", Font.PLAIN, 16));

        switch (a) {
            case "Archivo":
                b.add(agregarItemMenu("Nuevo archivo"));
                b.add(agregarItemMenu("Abrir archivo"));
                b.add(agregarItemMenu("Guardar"));
                b.add(agregarItemMenu("Guardar como"));
                b.add(agregarItemMenu("Salir"));
                break;
            case "Editar":
                b.add(agregarItemMenu("Deshacer"));
                b.add(agregarItemMenu("Cortar"));
                b.add(agregarItemMenu("Copiar"));
                b.add(agregarItemMenu("Pegar"));
                b.add(agregarItemMenu("Seleccionar todo"));
                break;
            case "Opciones":
                b.add(agregarItemMenu("Personalizar IDE"));
                b.add(agregarItemMenu("Ayuda"));
                b.add(agregarItemMenu("Acerca de"));
                break;
            case "Compilar":
                b.add(agregarItemMenu("Compilar"));
               
                
        }

        return b;
    }

    // Agrega cada JMenuItem al JMenu correspondiente del Menu
    // y tambien les agrega un ActionListener a cada boton.
    private static JMenuItem agregarItemMenu(String x) {
        JMenuItem item = new JMenuItem(x);
        item.setFont(new Font("TimesRoman", Font.PLAIN, 14));

        switch (x) {
            case "Nuevo archivo":
                item.setIcon(UIManager.getIcon("FileView.fileIcon"));
                item.addActionListener((ActionEvent e) -> {
                    if (codeArea.getText().trim().equals("")) codeArea.setText("");
                    else JOptionPane.showMessageDialog(null, "Guarda primero tu código");
                });
                break;
            case "Abrir archivo":
                item.setIcon(UIManager.getIcon("FileView.directoryIcon"));
                item.addActionListener((ActionEvent e) -> {
                    String nameFile = fa.seleccionarArchivo();
                    String textoCompleto ="";
                    System.out.println(nameFile);
                    File fichero = new File(nameFile);
                    Scanner s = null;

                    try {
                        // Leemos el contenido del fichero
                        System.out.println("... Leemos el contenido del fichero ...");
                        s = new Scanner(fichero);

                        // Leemos linea a linea el fichero
                        while (s.hasNextLine()) {
                            String linea = s.nextLine(); 	// Guardamos la linea en un String
                            String lineaA = textoCompleto;
                            textoCompleto=(lineaA + linea + "\n");
                        }
                    } catch (Exception ex) { System.out.println("Mensaje: " + ex.getMessage());
                    } finally {
                        // Cerramos el fichero tanto si la lectura ha sido correcta o no
                        try { if (s != null) s.close();
                        } catch (Exception ex2) { System.out.println("Mensaje 2: " + ex2.getMessage()); }
                    }
                    
                    System.out.println(textoCompleto);
                    codeArea.setText("");
                    paintText paint  = new paintText();
                    ArrayList<String> textoSeparado;
                    textoSeparado = paint.separarTexto( textoCompleto );
                    textoSeparado.stream().forEach((palabra) -> {
                        SimpleAttributeSet attrs = new SimpleAttributeSet();
                        if(paintText.checkString(palabra))StyleConstants.setForeground(attrs, Color.RED);
                        else StyleConstants.setForeground(attrs, Color.WHITE);
                        try {
                            codeArea.getStyledDocument().insertString(codeArea.getStyledDocument().getLength(), palabra, attrs);
                        } catch (BadLocationException ex) {
                            Logger.getLogger(CompiladorFrame.class.getName()).log(Level.SEVERE, null, ex);
                        }
                    });

                });
                
                break;
            case "Guardar":
                item.setIcon(UIManager.getIcon("FileView.floppyDriveIcon"));
                item.addActionListener((ActionEvent e) -> {
                    String nameFile = fa.seleccionarArchivo();
                    System.out.println(nameFile);
                    fa.guardarTxt(codeArea.getText(), nameFile);
                });
                break;
            case "Guardar como":
                item.setIcon(UIManager.getIcon("FileView.floppyDriveIcon"));
                item.addActionListener((ActionEvent e) -> {
                    String nameFile = fa.seleccionarCarpeta();
                    System.out.println(nameFile);
                    fa.guardarTxt(codeArea.getText(), nameFile);
                });
                break;
            case "Salir":
                item.addActionListener((ActionEvent e) -> {
                    int resp = JOptionPane.showConfirmDialog(null, "¿Desea cerrar?");
                    if (resp == 0) {
                        System.exit(0);
                    }
                });
                break;

            case "Deshacer":
                item.addActionListener((ActionEvent e) -> {
                    JOptionPane.showMessageDialog(null, "// codigo para deshacer");       /////// agregar codigo de ActionListener
                });
                break;
            case "Cortar":
                item.addActionListener((ActionEvent e) -> {
                    JOptionPane.showMessageDialog(null, "// codigo para cortar");         /////// agregar codigo de ActionListener
                });
                break;
            case "Copiar":
                item.addActionListener((ActionEvent e) -> {
                    JOptionPane.showMessageDialog(null, "// codigo para copiar");         /////// agregar codigo de ActionListener
                });
                break;
            case "Pegar":
                item.addActionListener((ActionEvent e) -> {
                    JOptionPane.showMessageDialog(null, "// codigo para pegar");          /////// agregar codigo de ActionListener
                });
                break;
            case "Seleccionar todo":
                item.addActionListener((ActionEvent e) -> {
                    JOptionPane.showMessageDialog(null, "// codigo para seleccionar todo");/////// agregar codigo de ActionListener
                });
                break;
            case "Personalizar IDE":
                item.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    JOptionPane.showMessageDialog(null, "// abrir nueva ventana que permita modificar (textsize, color, background, etc)");
                });
                break;
            case "Ayuda":
                item.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    JOptionPane.showMessageDialog(null, "// abrir nueva ventana que explique cosas");
                });
                break;
            case "Acerca de":
                item.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    JOptionPane.showMessageDialog(null, "// abrir nueva ventana que muestre informacion");
                });
                break;
            case "Compilar":
                 item.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    if(prerevisionCodigoPanel()){
                        ejecutadorGO x1 = new ejecutadorGO();
                        errorsArea.setText(x1.conexion());
                        lexicoTxt = x1.getOutputLexico();
                        sintacticoTxt = x1.getOutputSintac();
                        semanticoTxt = x1.getOutputSemant();
                        codigoInterTxt = x1.getOutputCodigo();
                        analizadorArea.setText(lexicoTxt);
                    }
                });
                
                 break;
        }
        return item;
    }

    // agrega el contenido del toolBar y lo retorna listo para agregarlo al panel
    private static JToolBar agregarBarraHerramientas() {
        ArrayList<String> palabras = new ArrayList<>();
        JToolBar barra = new JToolBar();
        barra.setBackground(Color.LIGHT_GRAY);
        barra.setFloatable(false);
        barra.setBorderPainted(true);
        

        //palabras.add("Compilar");
        palabras.add("Lexico");
        palabras.add("Sintactico");
        palabras.add("Semantico");
        palabras.add("Codigo Intermedio");
        palabras.stream().forEach((temp) -> {
            barra.add(agregarBotonBarra(temp));
        });

        return barra;
    }

    // agrega los botones correspondientes al toolbar con su ActionListener
    private static JButton agregarBotonBarra(String a) {
        JButton b = new JButton(a);
        b.setBorderPainted(true);
        b.setFont(new Font("Arial", Font.PLAIN, 20));

        switch (a) {
            case "Compilar":
                Icon icon = UIManager.getIcon("FileView.computerIcon");
                b.setIcon(icon);
                b.setForeground(Color.RED);
                b.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    if(prerevisionCodigoPanel()){
                        ejecutadorGO x = new ejecutadorGO();
                        errorsArea.setText(x.conexion());
                        lexicoTxt = x.getOutputLexico();
                        sintacticoTxt = x.getOutputSintac();
                        semanticoTxt = x.getOutputSemant();
                        codigoInterTxt = x.getOutputCodigo();
                        analizadorArea.setText(lexicoTxt);
                    }
                });
                break;
            case "Lexico":
                b.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    analizadorArea.setText(lexicoTxt);
                });
                break;
            case "Sintactico":
                b.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    analizadorArea.setText(sintacticoTxt);
                });
                break;
            case "Semantico":
                b.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    analizadorArea.setText(semanticoTxt);
                });
                break;
            case "Codigo Intermedio":
                b.addActionListener((ActionEvent e) -> {                                /////// agregar codigo de ActionListener
                    analizadorArea.setText(codigoInterTxt);
                });
                break;
        }

        return b;
    }
    
    /*////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    
    private static boolean checkString(String word){
        return palabrasReservadas.stream().anyMatch((PR) -> (word.equals(PR)));
    }
    
    public static void pintarArchivo(String text){
        String[] arrOfStr = text.split("\\s|\\n");
        System.out.println("-----------------------------------------------");
        System.out.println("longitud array de palabras: " + arrOfStr.length);
        codeArea.setText("");
        for(String a:arrOfStr){
            System.out.print("["+a+"]");
            insertStringArchivo(a);
        }
    }
    
    private static void insertString(String word){
                SimpleAttributeSet attrs = new SimpleAttributeSet();
                if(checkString(word))StyleConstants.setForeground(attrs, Color.BLUE);
                else StyleConstants.setForeground(attrs, Color.BLACK);
                try {
                    codeArea.getStyledDocument().remove(codeArea.getCaretPosition() - word.length(),word.length());
                    codeArea.getStyledDocument().insertString(codeArea.getStyledDocument().getLength(), word, attrs);
                    System.out.println("posicion cursor: " +codeArea.getCaretPosition());
                    System.out.println("posicion cursor - palabra: " +(codeArea.getCaretPosition()-word.length()));
                    System.out.println("longitud palabra: " +word.length());
                    System.out.println("palabra: " +word);
                } catch (BadLocationException ex) {
                    Logger.getLogger(CompiladorFrame.class.getName()).log(Level.SEVERE, null, ex);
                    System.out.println("posicion cursor: " +codeArea.getCaretPosition());
                    System.out.println("posicion cursor - palabra: " +(codeArea.getCaretPosition()-word.length()));
                    System.out.println("longitud palabra: " +word.length());
                    System.out.println("palabra: " +word);
                }
    }
    
    public static void palabrasReservadas(String texto){
        String[] arrOfStr = texto.split("\\s|\\n");
        System.out.println("-----------------------------------------------");
        System.out.println("longitud array de palabras: " + arrOfStr.length);
        for(String a:arrOfStr){
            System.out.print("["+a+"]");
        }
        insertString(arrOfStr[arrOfStr.length-1]);
    }
    
    private static void insertStringArchivo(String word){
                SimpleAttributeSet attrs = new SimpleAttributeSet();
                if(checkString(word))StyleConstants.setForeground(attrs, Color.BLUE);
                else StyleConstants.setForeground(attrs, Color.BLACK);
                try {
                    codeArea.getStyledDocument().insertString(codeArea.getStyledDocument().getLength(), word, attrs);
                    System.out.println("posicion cursor: " +codeArea.getCaretPosition());
                    System.out.println("posicion cursor - palabra: " +(codeArea.getCaretPosition()-word.length()));
                    System.out.println("longitud palabra: " +word.length());
                    System.out.println("palabra: " +word);
                } catch (BadLocationException ex) {
                    Logger.getLogger(CompiladorFrame.class.getName()).log(Level.SEVERE, null, ex);
                    System.out.println("posicion cursor: " +codeArea.getCaretPosition());
                    System.out.println("posicion cursor - palabra: " +(codeArea.getCaretPosition()-word.length()));
                    System.out.println("longitud palabra: " +word.length());
                    System.out.println("palabra: " +word);
                }
    }
    */////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    
    private static void limpiarArchivos(){
        fa.guardarTxt("", "auxiliar.txt");
    }
    
    private static boolean prerevisionCodigoPanel(){
        if(codeArea.getText().isEmpty()){ 
            errorsArea.setText("\n No hay codigo que procesar"); 
            return false;
        }/*
        String nameFile = fa.seleccionarCarpeta();
        fa.guardarTxt(codeArea.getText(), nameFile); // Hay un error aqui
        */
        fa.guardarTxt(codeArea.getText(), "auxiliar.txt");
        return true;
    }
    
}
