package GUI;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;

public class MyFrame extends JFrame implements ActionListener {

    /**
     *
     */
    private static final long serialVersionUID = 3438573214393962750L;

    // Fields
    JButton button;

    public MyFrame() {
        String title = "Java GUI";

        // Creating buttons
        this.button = new JButton("Clicky");
        this.button.setBounds(100, 100, 250, 100);
        this.button.addActionListener(e -> System.out.println("Clickw"));
        this.button.setFocusable(false);

        // Frame settings
        this.setTitle(title);
        this.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        // this.setResizable(false);
        this.setSize(500, 500);
        this.setLayout(null);
        // this.getContentPane().setBackground(new Color(50, 50, 50));

        // ImageIcon image = new ImageIcon("path");
        // frame.setIconImage(image.getImage());

        this.add(this.button);

        this.setVisible(true);
    }

    @Override
    public void actionPerformed(ActionEvent e) {
        if (e.getSource() == this.button) {
            System.out.println("Click");
        }
    }
}
