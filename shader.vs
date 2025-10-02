#version 330 core
// <vec2 position, vec2 texCoords>
layout (location = 0) in vec4 vertex; 
//layout (location = 1) in vec3 aColor;

out vec2 TexCoords;

uniform vec4 uvModel;
uniform mat4 model;
uniform mat4 projection;

void main()
{
    TexCoords = vertex.zw * uvModel.zw + uvModel.xy;
    gl_Position = projection * model * vec4(vertex.xy, 0.0, 1.0);
    //ourColor = aColor;
}
