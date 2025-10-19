#version 330 core

in vec2 TexCoords;
out vec4 color;

// Texture
uniform sampler2D image;

// Color
uniform vec4 ourColor;

void main()
{    
    color = texture(image, TexCoords) + ourColor;
}  
