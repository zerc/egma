#include <memory>

#include "window.hpp"

Window::Window(float width, float height, std::string title) {
    window = new sf::RenderWindow(sf::VideoMode(width, height), title);
    alive = true;
};


bool Window::is_alive() {
    if (!alive) {
        return false;
    }
    // check all the window's events that were triggered since the last iteration of the loop
    sf::Event event;
    while (window->pollEvent(event)) {
        // "close requested" event: we close the window
        if (event.type == sf::Event::Closed) {
            window->close();
            alive = false;
        }
    }

    return alive;
};


void Window::display() {
    window->resetGLStates();
    window->clear(sf::Color::Green);
    window->display();
};


std::shared_ptr<WindowInterface> create_window(float width, float height, std::string title) {
    return std::make_shared<Window>(width, height, title);
};
