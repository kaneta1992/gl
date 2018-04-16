#version 330

uniform sampler2D tex;
uniform sampler2D cherryTex;
uniform sampler2D kiraTex;
uniform sampler2D mask1;
uniform sampler2D mask3;
uniform sampler2D sunTex;
uniform sampler2D lensTex;

uniform float time;

in vec2 fragTexCoord;

out vec4 outputColor;

vec2 rotate(vec2 pos, float angle) {
    return vec2(cos(angle) * pos.x - sin(angle) * pos.y, sin(angle) * pos.x + cos(angle) * pos.y);
}

void main() {
    vec2 uv = fragTexCoord * 1;
    float mask3 = texture(mask3, uv).r;
    vec4 cardCol = texture(tex, uv + vec2(sin(time * 4.0 + uv.x * 5.0 + uv.y * 20.0), cos(time * 4.0 + uv.x * 20.0 + uv.y * 5.0)) * 0.005 * (1.0 - mask3));

    vec4 cherryCol = texture(cherryTex, uv * 1.0 + normalize(vec2(-0.5, -0.5)) * time * 0.25);
    vec4 cherryCol2 = texture(cherryTex, uv * 2.0 + normalize(vec2(-0.5, -0.5)) * time * 0.35);

    vec2 cherryRotateUV = uv - vec2(2.0, 0.0);
    vec4 cherryCol3 = texture(cherryTex, rotate(cherryRotateUV * 2.0, time* 0.1));

    vec4 kiraCol = texture(kiraTex, uv + vec2(time*0.01, 0.0));

    vec2 sunUV = uv + vec2(0.1, 0.1);
    vec4 sunCol = pow(texture(sunTex, clamp(rotate(sunUV, time* 0.2) * 0.5 + vec2(0.5, 0.5), 0.0, 1.0)) * 6.0, vec4(3.0));

    vec4 lensCol = texture(lensTex, uv);

    float mask1 = texture(mask1, uv).r;
    //mask1 = 1.0;
    vec3 resultCol = mix(cardCol.rgb, cherryCol.rgb, cherryCol.a * mask1);
    // resultCol = mix(resultCol.rgb, cherryCol2.rgb, cherryCol2.a * mask1);
    resultCol = mix(resultCol.rgb, cherryCol3.rgb, cherryCol3.a * mask1);
    // resultCol += kiraCol.rgb * (2.0 + sin(time * 10.0 + uv.x * 10.0)) * 0.5;
    resultCol += kiraCol.rgb * mask1 * (2.0 + sin(time * 5.0 + uv.x * 10.0)) * 0.5;
    resultCol += sunCol.rgb * mask1;
    resultCol += lensCol.rgb * (1.0 + (2.0 + sin(time * 25.0)) * 0.05);

    outputColor = vec4(resultCol, 1.0);
}