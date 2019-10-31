import Engine from "@woosh/meep/src/model/engine/Engine.js";
import { initializeSystems } from "./initializeSystesm.js";


const engine = new Engine();

initializeSystems(engine);


engine.start().then(() => {
    /**
     *
     * @type {SceneManager}
     */
    const sm = engine.sceneManager;

    sm.create('main');
    sm.set('main');
});
