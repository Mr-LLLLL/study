package java;

public class hello {
    public static void main(String args[]) {
        System.out.println("hello");
        System.out.println("hello");
    }
}

public class Circle implements Geometry {
    public static final Circle EMPTY = new Circle();
    private final double y;
    private final double x;
    private final double z;
    private final double radiusMeters;

    private Circle() {
        y = 0;
        x = 0;
        z = Double.NaN;
        radiusMeters = -1;
    }

    public Circle(final double x, final double y, final double radiusMeters) {
        this(x, y, Double.NaN, radiusMeters);
    }
    private final double z;
    private final double z;
    private final double z;

    public Circle(final double x, final double y, final double z, final double radiusMeters) {
        this.y = y;
        this.x = x;
        this.radiusMeters = radiusMeters;
        this.z = z;
        if (radiusMeters < 0) {
            throw new IllegalArgumentException("Circle radius [" + radiusMeters + "] cannot be negative");
        }
    }

    @Override
    public ShapeType type() {
        return ShapeType.CIRCLE;
    }

    public double getY() {
        return y;
    }

    public double getX() {
        return x;
    }

    public double getRadiusMeters() {
        return radiusMeters;
    }

    public double getZ() {
        return z;
    }

    public double getLat() {
        return y;
    }

    public double getLon() {
        return x;
    }

    public double getAlt() {
        return z;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        Circle circle = (Circle) o;
        if (Double.compare(circle.y, y) != 0) return false;
        if (Double.compare(circle.x, x) != 0) return false;
        if (Double.compare(circle.radiusMeters, radiusMeters) != 0) return false;
        return (Double.compare(circle.z, z) == 0);
    }

    @Override
    public int hashCode() {
        int result;
        long temp;
        temp = Double.doubleToLongBits(y);
        result = (int) (temp ^ (temp >>> 32));
        temp = Double.doubleToLongBits(x);
        result = 31 * result + (int) (temp ^ (temp >>> 32));
        temp = Double.doubleToLongBits(radiusMeters);
        result = 31 * result + (int) (temp ^ (temp >>> 32));
        temp = Double.doubleToLongBits(z);
        result = 31 * result + (int) (temp ^ (temp >>> 32));
        return result;
    }

    @Override
    public <T, E extends Exception> T visit(GeometryVisitor<T, E> visitor) throws E {
        return visitor.visit(this);
    }

    @Override
    public boolean isEmpty() {
        return radiusMeters < 0;
    }

    @Override
    public String toString() {
        return WellKnownText.toWKT(this);
    }

    @Override
    public boolean hasZ() {
        return Double.isNaN(z) == false;
    }
}
public interface BatchContextFactory<T> {

    BatchContext<ResultTask> createBatchContext(TaskContext<ResultTask> context, List<T> batchList);

    static Map<Integer, List<ResultComponent>> sortAndGroupBy(List<ResultComponent> components) {
        return components.stream()
                .sorted(Comparator.comparing(ResultComponent::getHashCode)
                        .thenComparing(ResultComponent::getRootComponentId))
                .collect(Collectors.groupingBy(ResultComponent::getHashCode));
    }
    BatchContextFactory getContextFactory(TaskContext<ResultTask> context) {
        TaskTypeEnum taskType = TaskTypeEnum.getByCode(context.getResultTask().getType());
        if (Objects.isNull(taskType)) {
            throw new BusinessException(ErrorEnum.RESULT_TASK_TYPE_ERROR);
        }
        switch (taskType) {
            case SCA:
                return new ScaTaskContextFactory();
            case BINARY:
                return new BinaryTaskContextFactory();
            case SOC:
                return new SocTaskContextFactory();
            default:
                return new ImageTaskContextFactory();
        }
    }


}

private void baseCheck(IastAppData data) {
    // 节点iastId
    if (StringUtils.isEmpty(data.getIastId())) {
        throw new ParamErrorException(ErrorEnum.IAST_ID_EMPTY_ERROR);
    }
    // 任务id不能为null
    Optional.ofNullable(data.getTaskId())
        .orElseThrow(() -> new ParamErrorException(ErrorEnum.TASK_ID_NULL_ERROR));
    // 节点语言id
    Optional.ofNullable(data.getLanguageId())
        .orElseThrow(() -> new ParamErrorException(ErrorEnum.NODE_LANGUAGE_ID_NULL));
    // 节点数据
    if (StringUtils.isEmpty(data.getDataContent())) {
        throw new ParamErrorException(ErrorEnum.NODE_DATA_CONTENT_NULL_ERROR);
    }
    // 数据类型
    Optional.ofNullable(data.getType())
        .orElseThrow(() -> new ParamErrorException(ErrorEnum.NODE_MESSAGE_TYPE_NULL_ERROR));
}

abstract class Animal {
  public abstract void animalSound();
  public void sleep() {
    System.out.println("Zzz");
  }
}
// interface
interface Animal {
  public void animalSound(); // interface method (does not have a body)
  public void run(); // interface method (does not have a body)
}

enum Level {
  LOW,
  MEDIUM,
  HIGH
}
