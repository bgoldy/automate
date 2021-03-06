import { Component, Event, EventEmitter, Prop } from '@stencil/core';

/**
 * @description
 * chef-data-point is used in conjunction with various charts to specify data points.
 *
 * @example
 * <style>
 *   chef-radial-chart .failed, chef-radial-chart .critical {
 *     color: hsl(var(--chef-critical));
 *   }
 *   chef-radial-chart .success, chef-radial-chart .major {
 *     color: hsl(var(--chef-success));
 *   }
 *   chef-radial-chart .skipped, chef-radial-chart .minor {
 *     color: hsl(var(--chef-grey));
 *   }
 * </style>
 * <chef-radial-chart style="width: 220px; height: 220px;">
 *   <span slot="innerText">Text for the center of the chart</span>
 *
 *   <chef-data-point value="4" class="failed">4 Failed</chef-data-point>
 *   <chef-data-point value="3" class="success">3 Successful</chef-data-point>
 *   <chef-data-point value="2" class="skipped">2 Skipped</chef-data-point>
 *
 *   <chef-data-point value="3" secondary class="critical">Critical</chef-data-point>
 *   <chef-data-point value="2" secondary class="major">Major</chef-data-point>
 *   <chef-data-point value="1" secondary class="minor">Minor</chef-data-point>
 * </chef-radial-chart>
 */
@Component({
  tag: 'chef-data-point'
})
export class ChefDataPoint {

  /**
   * The value assigned to the data point.
   */
  @Prop() value = 0;

  /**
   * Some charts support special secondary data points.
   */
  @Prop() secondary = false;

  @Event() updated: EventEmitter;

  componentWillUpdate() {
    this.updated.emit();
  }

  render() {
    return (
      <slot />
    );
  }

}
