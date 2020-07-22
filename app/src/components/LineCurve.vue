<template>
  <canvas class="line-curve"></canvas>
</template>

<script>
export default {
  props: {
    x1: { type: Number, required: true },
    y1: { type: Number, required: true },
    x2: { type: Number, required: true },
    y2: { type: Number, required: true },
    xWeight1: { type: Number, required: true },
    xWeight2: { type: Number, required: true },
  },

  mounted() {
    const canvas = this.$el;
    if (canvas != undefined) {
      /*I don't think canvas would be undefined (now), but I used created() before.*/
      const padding = 5;
      const widen = 2;
      // Update the bounds of the canvas, with additional padding.
      canvas.width = this.x2 - this.x1 + padding * 2;
      canvas.height = Math.abs(this.y2 - this.y1) + padding * 2;
      canvas.style.left = this.x1 - 5 + "px";
      canvas.style.top = Math.min(this.y1, this.y2) - 5 + "px";

      // Setup the canvas context & drawing style - just going for a predefined style here, as we only need it for 1 view.
      const ctx = canvas.getContext("2d");
      ctx.strokeStyle = "#303133";
      ctx.lineWidth = 1;

      // Only for the y coordinate we might have y2 be less than y1. Expecting all curves to go left to right.
      const minY = Math.min(this.y1, this.y2);

      // Draws the curve with 2 control points adjusted between x1 and x2 by xWeight1 and xWeight2 (0 = at x1, 1 = at x2).
      ctx.moveTo(padding - widen, padding + this.y1 - minY);
      ctx.bezierCurveTo(
        padding + (this.x2 - this.x1) * this.xWeight1,
        padding + this.y1 - minY,
        padding + (this.x2 - this.x1) * this.xWeight2,
        padding + this.y2 - minY,
        padding + this.x2 - this.x1 + widen,
        padding + this.y2 - minY
      );
      ctx.stroke();
    }
  },
};
</script>

<style lang="scss" scoped>
.line-curve {
  position: absolute;
}
</style>
