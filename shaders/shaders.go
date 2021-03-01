package shaders

const (
	VertexShader = `
		#version 410
		layout(location = 0) in vec3 vertexPos;
		uniform mat4 projection;
		uniform mat4 camera;
		uniform mat4 model;
		void main()	{
			gl_Position = projection * camera * model * vec4(vertexPos, 1.0);
		}
	` + "\x00"

	FragmentShader = `
		#version 410
		out vec3 color;
		void main() {
			color = vec3(1, 0, 0);
		}
	` + "\x00"
)