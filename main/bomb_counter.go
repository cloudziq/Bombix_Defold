components {
  id: "bomb_counter"
  component: "/main/bomb_counter.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "scale {\n"
  "  x: 0.64\n"
  "  y: 0.64\n"
  "  z: 1.0\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.6\n"
  "  y: 0.6\n"
  "  z: 0.6\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.101960786\n"
  "  y: 0.2\n"
  "  z: 0.6\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"X\"\n"
  "font: \"/assets/fonts/system.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 32.0
    y: 30.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
