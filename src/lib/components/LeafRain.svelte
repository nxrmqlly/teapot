<script lang="ts">
    import Leaf from "./Leaf.svelte";

    interface Props {
        density?: number;
    }

    const { density = 24 } = $props();

    interface LeafParticle {
        id: number;
        variant: 1 | 2 | 3 | 4;
        left: number;
        size: number;
        duration: number;
        delay: number;
        drift: number;
        rotation: number;
        opacity: number;
    }

    const leaves: LeafParticle[] = $derived(
        Array.from({ length: density }, (_, id) => ({
            id,
            variant: (Math.floor(Math.random() * 4) + 1) as 1 | 2 | 3 | 4,
            left: Math.random() * 100,
            size: 18 + Math.random() * 25,
            duration: 9 + Math.random() * 8,
            delay: Math.random() * -17,
            drift: -120 + Math.random() * 240,
            rotation: -360 + Math.random() * 720,
            opacity: 0.3 + Math.random() * 0.45,
        })),
    );
</script>

<div class="pointer-events-none fixed inset-0 z-0 overflow-hidden">
    {#each leaves as leaf}
        <div
            class="leaf absolute"
            style="
                left: {leaf.left}%;
                width: {leaf.size}px;
                height: {leaf.size}px;
                animation-duration: {leaf.duration}s;
                animation-delay: {leaf.delay}s;
                opacity: {leaf.opacity};
                --drift: {leaf.drift}px;
                --rotation: {leaf.rotation}deg;
            "
        >
            <Leaf variant={leaf.variant} classes="h-full w-full" />
        </div>
    {/each}
</div>

<style>
    .leaf {
        top: -60px;
        animation: leaf-fall linear infinite;
        will-change: transform;
    }

    @keyframes leaf-fall {
        0% {
            transform: translate3d(0, -60px, 0) rotate(0deg);
        }

        20% {
            transform: translate3d(calc(var(--drift) * -0.15), 20vh, 0)
                rotate(calc(var(--rotation) * 0.2));
        }

        45% {
            transform: translate3d(calc(var(--drift) * 0.5), 45vh, 0)
                rotate(calc(var(--rotation) * 0.5));
        }

        70% {
            transform: translate3d(calc(var(--drift) * 0.15), 70vh, 0)
                rotate(calc(var(--rotation) * 0.75));
        }

        100% {
            transform: translate3d(var(--drift), calc(100vh + 80px), 0)
                rotate(var(--rotation));
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .leaf {
            animation: none;
            display: none;
        }
    }
</style>
