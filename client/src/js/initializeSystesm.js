import GUIElementSystem from "@woosh/meep/src/model/engine/ecs/systems/GUIElementSystem.js";
import ScriptSystem from "@woosh/meep/src/model/engine/ecs/systems/ScriptSystem.js";
import TeamSystem from "@woosh/meep/src/model/engine/ecs/team/TeamSystem.js";
import PathFollowingSystem from "@woosh/meep/src/model/navigation/ecs/systems/PathFollowingSystem.js";
import PathSystem from "@woosh/meep/src/model/navigation/ecs/systems/PathSystem.js";
import SteeringSystem from "@woosh/meep/src/model/engine/ecs/systems/SteeringSystem.js";
import MotionSystem from "@woosh/meep/src/model/engine/ecs/systems/MotionSystem.js";
import TagSystem from "@woosh/meep/src/model/engine/ecs/systems/TagSystem.js";
import { ParticleEmitterSystem2 } from "@woosh/meep/src/model/graphics/particles/ecs/ParticleEmitterSystem2.js";
import { SoundEmitterSystem } from "@woosh/meep/src/model/sound/ecs/SoundEmitterSystem.js";
import SoundControllerSystem from "@woosh/meep/src/model/sound/ecs/SoundControllerSystem.js";
import SoundListenerSystem from "@woosh/meep/src/model/sound/ecs/SoundListenerSystem.js";
import MortalitySystem from "@woosh/meep/src/model/engine/ecs/systems/MortalitySystem.js";
import TimerSystem from "@woosh/meep/src/model/engine/ecs/systems/TimerSystem.js";
import HeadsUpDisplaySystem from "@woosh/meep/src/model/engine/ecs/systems/HeadsUpDisplaySystem.js";
import AnimationSystem from "@woosh/meep/src/model/engine/ecs/systems/AnimationSystem.js";
import TopDownCameraControllerSystem
    from "@woosh/meep/src/model/engine/input/ecs/systems/TopDownCameraControllerSystem.js";
import TransformSystem from "@woosh/meep/src/model/engine/ecs/systems/TransformSystem.js";
import { TopDownCameraLanderSystem } from "@woosh/meep/src/model/engine/input/ecs/systems/TopDownCameraLanderSystem.js";
import { CameraSystem } from "@woosh/meep/src/model/graphics/ecs/camera/CameraSystem.js";
import { MeshSystem } from "@woosh/meep/src/model/graphics/ecs/mesh/MeshSystem.js";
import ClingToTerrainSystem from "@woosh/meep/src/model/level/terrain/ecs/ClingToTerrainSystem.js";
import TerrainSystem from "@woosh/meep/src/model/level/terrain/ecs/TerrainSystem.js";
import WaterSystem from "@woosh/meep/src/model/graphics/ecs/water/WaterSystem.js";
import TrailSystem from "@woosh/meep/src/model/graphics/ecs/trail/TrailSystem.js";
import Trail2DSystem from "@woosh/meep/src/model/graphics/ecs/trail2d/Trail2DSystem.js";
import { Foliage2System } from "@woosh/meep/src/model/level/foliage/ecs/Foliage2System.js";
import ViewportPositionSystem from "@woosh/meep/src/model/engine/ecs/systems/ViewportPositionSystem.js";
import { GridPosition2TransformSystem } from "@woosh/meep/src/model/engine/grid/systems/GridPosition2TransformSystem.js";
import SynchronizePositionSystem from "@woosh/meep/src/model/engine/ecs/systems/SynchronizePositionSystem.js";
import GridObstacleSystem from "@woosh/meep/src/model/engine/grid/systems/GridObstacleSystem.js";
import GridPositionSystem from "@woosh/meep/src/model/engine/grid/systems/GridPositionSystem.js";
import { InputSystem } from "@woosh/meep/src/model/engine/input/ecs/systems/InputSystem.js";
import HighlightSystem from "@woosh/meep/src/model/graphics/ecs/highlight/HighlightSystem.js";
import LightSystem from "@woosh/meep/src/model/graphics/ecs/light/LightSystem.js";
import AnimationControllerSystem from "@woosh/meep/src/model/graphics/ecs/animation/AnimationControllerSystem.js";
import PropertySetSystem from "@woosh/meep/src/model/engine/ecs/systems/PropertySetSystem.js";
import { FogOfWarSystem } from "@woosh/meep/src/model/level/fow/FogOfWarSystem.js";
import { FogOfWarRevealerSystem } from "@woosh/meep/src/model/level/fow/FogOfWarRevealerSystem.js";
import { BlackboardSystem } from "@woosh/meep/src/model/engine/intelligence/blackboard/BlackboardSystem.js";
import { BehaviorSystem } from "@woosh/meep/src/model/engine/intelligence/behavior/ecs/BehaviorSystem.js";
import { SerializationMetadataSystem } from "@woosh/meep/src/model/engine/ecs/systems/SerializationMetadataSystem.js";
import { AttachmentSocketsSystem } from "@woosh/meep/src/model/engine/ecs/sockets/AttachmentSocketsSystem.js";
import { AttachmentSystem } from "@woosh/meep/src/model/engine/ecs/attachment/AttachmentSystem.js";

/**
 *
 * @param {Engine} engine
 */
function initializeSystems(
    engine
) {

    const entityManager = engine.entityManager;
    const graphics = engine.graphics;
    const sound = engine.sound;
    const assetManager = engine.assetManager;
    const grid = engine.grid;
    const devices = engine.devices;

    const guiSystem = new GUIElementSystem(engine.gui.view);
    const headsUpDisplaySystem = new HeadsUpDisplaySystem(graphics);

    entityManager
        .addSystem(new ScriptSystem())
        .addSystem(new TeamSystem())
        .addSystem(new PathFollowingSystem())
        .addSystem(new PathSystem())
        .addSystem(new SteeringSystem())
        .addSystem(new MotionSystem())
        .addSystem(new TagSystem())
        .addSystem(new ParticleEmitterSystem2(assetManager, graphics))
        .addSystem(new SoundEmitterSystem(assetManager, sound.destination, sound.context))
        .addSystem(new SoundControllerSystem())
        .addSystem(new SoundListenerSystem(sound.context))
        .addSystem(new MortalitySystem())
        .addSystem(new TimerSystem())
        .addSystem(guiSystem)
        .addSystem(new TransformSystem())
        .addSystem(new AnimationSystem(graphics.viewport.size))
        .addSystem(new TopDownCameraControllerSystem(graphics))
        .addSystem(new TopDownCameraLanderSystem())
        .addSystem(new CameraSystem(graphics.scene, graphics))
        .addSystem(new MeshSystem(graphics, assetManager))
        .addSystem(new ClingToTerrainSystem())
        .addSystem(new TerrainSystem(graphics, grid, assetManager))
        .addSystem(new WaterSystem(graphics))
        .addSystem(new TrailSystem(graphics))
        .addSystem(new Trail2DSystem(graphics, assetManager))
        .addSystem(new Foliage2System(assetManager, graphics))
        .addSystem(new ViewportPositionSystem(graphics.viewport.size))
        .addSystem(new GridPosition2TransformSystem())
        .addSystem(new SynchronizePositionSystem())
        .addSystem(new GridObstacleSystem(grid))
        .addSystem(new GridPositionSystem())
        .addSystem(new InputSystem(devices))
        .addSystem(new HighlightSystem(graphics))
        .addSystem(new LightSystem(graphics.scene, {
            shadowResolution: 1024
        }))
        .addSystem(new AnimationControllerSystem())
        .addSystem(new PropertySetSystem())
        .addSystem(headsUpDisplaySystem)
        .addSystem(new FogOfWarSystem(graphics))
        .addSystem(new FogOfWarRevealerSystem(0))
        .addSystem(new BlackboardSystem())
        .addSystem(new BehaviorSystem())
        .addSystem(new SerializationMetadataSystem())
        .addSystem(new AttachmentSocketsSystem())
        .addSystem(new AttachmentSystem())
    ;
}

export {
    initializeSystems
};
