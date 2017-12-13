CASC="/home/mjibson/src/github.com/nydus/build/bin/storm-extract -i hots"

${CASC} -t txt -x -v
${CASC} -t xml -x -v

${CASC} -t dds -s mods/heroes.stormmod/base.stormassets/Assets/Textures -f storm_btn_d3 -x -v
${CASC} -t dds -s mods/heroes.stormmod/base.stormassets/Assets/Textures -f storm_ui_ingame_ -x -v
${CASC} -t dds -s mods/heroes.stormmod/base.stormassets/Assets/Textures -f storm_temp_ -x -v
${CASC} -t dds -s mods/heroes.stormmod/base.stormassets/Assets/Textures -f storm_ui_icon -x -v
