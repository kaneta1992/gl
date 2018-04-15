#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 position;
in vec2 uv;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = uv;
    gl_Position = projection * camera * model * vec4(position, 1);
}