
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
    vec4 cardCol = texture(tex, uv + vec2(sin(time * 3.0 + uv.x * 5.0 + uv.y * 20.0), cos(time * 3.0 + uv.x * 20.0 + uv.y * 5.0)) * 0.005 * (1.0 - mask3));

    vec4 cherryCol = texture(cherryTex, uv * 0.75 + normalize(vec2(-0.0, 1.0)) * time * 0.025);

    vec2 cherryRotateUV = uv - vec2(2.0, 0.0);
    vec4 cherryCol3 = texture(cherryTex, rotate(cherryRotateUV * 1.0, time* 0.01));

    vec4 kiraCol = texture(kiraTex, uv);

    vec4 sunCol = texture(sunTex, uv);

    vec4 lensCol = texture(lensTex, uv);

    float mask1 = texture(mask1, uv).r;
    //mask1 = 1.0;
    vec3 resultCol = cardCol.rgb;
    //resultCol = mix(resultCol.rgb, cherryCol.rgb, cherryCol.a * mask1);
    // resultCol = mix(resultCol.rgb, cherryCol2.rgb, cherryCol2.a * mask1);
    //resultCol = mix(resultCol.rgb, cherryCol3.rgb, cherryCol3.a * mask1);
    // resultCol += kiraCol.rgb * (2.0 + sin(time * 10.0 + uv.x * 10.0)) * 0.5;
    resultCol += cherryCol.rgb * mask1 * (1.0 + (2.0 + sin(time * 15.0 + uv.x * 10.0)) * 0.2);
    resultCol += cherryCol3.rgb * mask1 * (1.0 + (2.0 + sin(time * 10.0 + uv.x * 10.0)) * 0.2);
    resultCol += sunCol.rgb * (0.5 + (2.0 + sin(time * 5.0)) * 0.075);
    resultCol += kiraCol.rgb * (0.0 + (2.0 + sin(time * 5.0+ uv.x * 10.0+ uv.y * 10.0)) * 0.5);

    outputColor = vec4(resultCol, 1.0);
}
