package shaders

const (
	VertexShader = `
		#version 410
		layout(location = 0) in vec3 vertexPos;
		layout(location = 1) in vec3 vertexColor;
		uniform mat4 projection;
		uniform mat4 camera;
		uniform mat4 model;
		out vec3 fragmentColor;
		void main()	{
			gl_Position = projection * camera * model * vec4(vertexPos, 1.0);
			fragmentColor = vertexColor;
		}
	` + "\x00"

	FragmentShader = `
		#version 410
		in vec3 fragmentColor;
		out vec3 color;
		void main() {
			color = fragmentColor;
		}
	` + "\x00"
)