#version 330 core

in vec4 in_vert;

out vec4 vert;

void main() {
    gl_Position = in_vert;
    vert = in_vert;
}
