#include <SFML/Window.hpp>
#include <SFML/Graphics.hpp>

#include "config.cpp"

void create_window(Config& config) {
    sf::RenderWindow window(sf::VideoMode(config.scene.width, config.scene.height), "Test");

    // run the program as long as the window is open
    while (window.isOpen())
    {
        // check all the window's events that were triggered since the last iteration of the loop
        sf::Event event;
        while (window.pollEvent(event))
        {
            // "close requested" event: we close the window
            if (event.type == sf::Event::Closed)
                window.close();
        }

        // 
        window.resetGLStates();
        window.clear(sf::Color::Green);
        window.display();
    }
}