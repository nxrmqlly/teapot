<script lang="ts">
    const { title, children } = $props();
</script>

<div class="parent bg-tea-olive rounded-4xl w-full h-full">
    <div
        class="inner border-2 border-tea-olive bg-tea-yellow-green rounded-4xl w-full h-full"
    >
        <div class="content m-5">
            <h3 class="font-inst-sans italic text-[32px] leading-7">
                {title}
            </h3>
            <div class="mt-4">
                {@render children()}
            </div>
        </div>
    </div>
</div>

<style>
    .inner {
        transform: translateY(-0.25em) translateX(-0.25em);
    }
</style>
