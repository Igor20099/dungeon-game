components {
  id: "enemy"
  component: "/main/assets/gui/enemies/orc.gui"
}
components {
  id: "death"
  component: "/main/assets/scripts/components/death.script"
}
components {
  id: "enemy_stats"
  component: "/main/assets/scripts/enemy/enemy_stats.script"
  properties {
    id: "type"
    value: "orc"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "health"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "damage"
    value: "20.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "delay_attack"
    value: "3.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "coins"
    value: "300.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "exp"
    value: "200.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "name"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"\\320\\236\\321\\200\\320\\272\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    x: 614.0
    y: 128.0
  }
  scale {
    x: 1.0E-6
    y: 1.0E-6
  }
}
