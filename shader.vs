#version 330 core
// Split into vec2 position, vec2 texCoords
layout (location = 0) in vec4 vertex; 

out vec2 TexCoords;

// Source
uniform vec4 sourceUV;

// Destination
uniform mat4 model;

// Projection of the window
uniform mat4 projection;

void main()
{
    // sourceUV.xy is the offset
    // sourceUV.zw is the size
    TexCoords = vertex.zw * sourceUV.zw + sourceUV.xy;

    gl_Position = projection * model * vec4(vertex.xy, 0.0, 1.0);
}
