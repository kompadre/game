import { BinarySerializationRegistry } from "@woosh/meep/src/model/engine/ecs/storage/binary/BinarySerializationRegistry.js";
import { TransformSerializationAdapter } from "@woosh/meep/src/model/engine/ecs/components/Transform.js";
import { TeamSerializationAdapter } from "@woosh/meep/src/model/engine/ecs/team/Team.js";

export const gameBinarySerializationRegistry = new BinarySerializationRegistry();


gameBinarySerializationRegistry.registerAdapters([
    new TransformSerializationAdapter(),
    new MinimapMarkerSerializationAdapter(),
    new NameSerializationAdapter(),
    new TeamSerializationAdapter(),
    new ItemContainerSerializationAdapter(),
    new PathFollowerSerializationAdapter(),
    new PathSerializationAdapter(),
    new TagSerializationAdapter(),
    new ParticleEmitterSerializationAdapter(),
    new SoundEmitterSerializationAdapter(),
    new SoundControllerSerializationAdapter(),
    new SoundListenerSerializationAdapter(),
    new AnimationSerializationAdapter(),
    new MeshSerializationAdapter(),
    new TopDownCameraControllerSerializationAdapter(),
    new TopDownCameraLanderSerializationAdapter(),
    new CameraSerializationAdapter(),
    new ClingToTerrainSerializationAdapter(),
    new TerrainSerializationAdapter(),
    new WaterSerializationAdapter(),
    new InstancedMeshSerializationAdapter(),
    new GridPosition2TransformSerializationAdapter(),
    new GridObstacleSerializationAdapter(),
    new GridPositionSerializationAdapter(),
    new HighlightSerializationAdapter(),
    new LightSerializationAdapter(),
    new AnimationControllerSerializationAdapter(),
    new PropertySetSerializationAdapter(),
    new FogOfWarSerializationAdapter(),
    new FogOfWarRevealerSerializationAdapter(),
    new BlackboardSerializationAdapter(),
    new SerializationMetadataSerializationAdapter(),
    new QuesterSerializationAdapter(),
    new FacingDirectionSerializationAdapter(),
    new SteeringSerializationAdapter(),
    new MotionSerializationAdapter(),
    new CombatUnitSerializationAdapter(),
]);
