use utils;

#[derive(Clone)]
struct Light {
    x: i32,
    y: i32,
    lit: bool,
}

#[derive(Clone)]
struct LightGrid {
    width: i32,
    height: i32, 
    lights: Vec<Light>,
}

impl LightGrid {
    // maximum possible size of the grid
    pub fn max_size(&self) -> i32 {
        self.width * self.height
    }

    pub fn add(&mut self, i: Light) {
        &self.lights.push(i);
    } 

    // returns the current size of the grid, ie number of elements in there
    pub fn size(&self) -> usize {
        self.lights.len()
    }
}

fn fill_grid(grid: &LightGrid) {

}

#[test]
fn will_fill_the_grid_with_lights_in_default_state(){
    let width = 5;
    let height = 5;
    let mut grid = LightGrid{ width, height, lights: Vec::new() };
    let mut i: i32 = 0;
    let size = grid.max_size();
    while i < size {
        let x = i % width;
        // integer division for y pos
        let y = i / width;
        &grid.add(Light{ x, y, lit: false });
        i += 1;
        assert_eq!(i, grid.size() as i32);
    }
}

fn calculate_lights(data: &str) -> i32 {
    0
}

pub fn solve() -> String {
  let data = utils::read_file("../data/2015_6.txt");
  let lights_lit = calculate_lights(&data);
  format!(
    "{:?} christmas lights have been lit",
    &lights_lit
  )
}
